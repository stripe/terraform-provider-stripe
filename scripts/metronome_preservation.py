#!/usr/bin/env python3
"""Freeze and check complete provider artifacts without transforming their bytes."""

import argparse
import hashlib
import json
import os
from pathlib import Path, PurePosixPath
import re
import subprocess
import sys


SHARED = {
    "main.go": "Select the combined provider entry point.",
    "go.mod": "Declare integration dependencies.",
    "go.sum": "Lock integration dependency checksums.",
    "docs/index.md": "Document combined provider configuration.",
}
CHANNELS = {"ga", "beta", "private-preview"}
REVISION = re.compile(r"^[0-9a-f]{40}$")
SHA256 = re.compile(r"^[0-9a-f]{64}$")
METRONOME_INPUTS = {
    f"packages/terraform/src/metronome/api/{name}.json" for name in (
        "credential-mapping-exceptions", "credential-operations",
        "lifecycle-schema-policy", "managed-resources", "managed-schema-policy",
        "metronome-openapi", "operation-inventory.generated", "read-operations",
        "security-review", "terraform-mapping",
    )
} | {
    "packages/terraform/src/metronome/sdk-bindings.json",
    "packages/terraform/src/metronome/sdk-endpoints.json",
    "packages/terraform/src/metronome/LICENSE",
    "packages/terraform/src/metronome/SDK-LICENSE",
}


def git(repo, *args, data=None):
    return subprocess.run(
        ["git", "-C", str(repo), *args], input=data, check=True,
        stdout=subprocess.PIPE, stderr=subprocess.PIPE,
    ).stdout


def relative_path(value):
    if not isinstance(value, str) or not value or "\\" in value:
        raise ValueError(f"Invalid relative path: {value!r}")
    path = PurePosixPath(value)
    if path.is_absolute() or str(path) != value or any(
        part in {".", "..", ".git"} for part in path.parts
    ) or any(ord(c) < 32 for c in value):
        raise ValueError(f"Invalid relative path: {value!r}")
    return value


def metadata(path, content):
    return {"path": path, "length": len(content),
            "sha256": hashlib.sha256(content).hexdigest()}


def git_files(repo, revision):
    if not isinstance(revision, str) or not REVISION.fullmatch(revision):
        raise ValueError("A full immutable Git commit is required")
    commit = git(repo, "rev-parse", f"{revision}^{{commit}}").decode().strip()
    if commit != revision:
        raise ValueError("Provider revision is not a commit")
    entries = git(repo, "ls-tree", "-rz", revision).split(b"\0")
    paths, objects = [], []
    for entry in entries:
        if not entry:
            continue
        header, raw_path = entry.split(b"\t", 1)
        mode, kind, oid = header.split()
        if kind != b"blob" or mode not in {b"100644", b"100755"}:
            raise ValueError(f"Unsupported baseline file mode: {entry!r}")
        paths.append(relative_path(raw_path.decode("utf-8")))
        objects.append(oid)
    if not paths:
        raise ValueError("Empty baseline inventory")
    response = git(repo, "cat-file", "--batch", data=b"\n".join(objects) + b"\n")
    files, offset = {}, 0
    for path in paths:
        end = response.index(b"\n", offset)
        size = int(response[offset:end].split()[2])
        files[path] = response[end + 1:end + 1 + size]
        offset = end + 2 + size
    return files


def validate_additions(additions, existing):
    seen = set()
    for entry in additions:
        path = relative_path(entry["path"])
        owner = entry["owner"]
        if path in existing or path in seen or not entry.get("purpose", "").strip():
            raise ValueError(f"Duplicate, existing or unexplained addition: {path}")
        seen.add(path)
        parts = PurePosixPath(path).parts
        metronome = any("metronome" in part for part in parts)
        allowed = (
            owner == "metronome" and metronome and parts[0] != "stripe-go"
            or owner == "composition" and path.startswith((
                "internal/combinedprovider/", "internal/combined_provider/",
                "internal/composition/", "internal/providerdata/",
            ))
            or owner == "shared-runtime" and path.startswith("internal/terraformruntime/") and path.endswith(".go")
            or owner == "provenance" and path in {
                "generation-provenance.json", "stripe-preservation-manifest.json",
            }
            or owner == "ci" and path.startswith(".github/workflows/")
        )
        if not allowed:
            raise ValueError(f"Unexpected Stripe-owned addition: {path} ({owner})")


def validate_manifest(manifest, originals):
    if manifest.get("schemaVersion") != 1 or manifest.get("channel") not in CHANNELS:
        raise ValueError("Unsupported preservation manifest schema or channel")
    protected = manifest.get("protected", [])
    if not protected:
        raise ValueError("Empty protected comparison set")
    exceptions = manifest.get("sharedIntegrationFiles", [])
    if {e["path"]: e["purpose"] for e in exceptions} != SHARED:
        raise ValueError("Shared integration exceptions must be the four exact approved paths")
    recorded = {}
    for entry in protected + exceptions:
        path = relative_path(entry["path"])
        if path in recorded:
            raise ValueError(f"Duplicate baseline path: {path}")
        if type(entry["length"]) is not int or entry["length"] < 0 or not SHA256.fullmatch(entry["sha256"]):
            raise ValueError(f"Invalid length or SHA256: {path}")
        recorded[path] = entry
    if set(recorded) != set(originals):
        raise ValueError("Manifest does not cover the complete frozen provider tree")
    for path, content in originals.items():
        expected = metadata(path, content)
        if any(recorded[path][key] != expected[key] for key in expected):
            raise ValueError(f"Frozen baseline metadata differs from Git bytes: {path}")
    if set(e["path"] for e in protected) != set(originals) - set(SHARED):
        raise ValueError("Every existing file outside the four exceptions must be protected")
    validate_additions(manifest.get("allowedAdditions", []), recorded)


def artifact_files(root):
    root = Path(root)
    if root.is_symlink() or not root.is_dir():
        raise ValueError("Artifact must be a real directory")
    files = {}
    for directory, dirs, names in os.walk(root, followlinks=False):
        for name in dirs[:]:
            path = Path(directory) / name
            if path.is_symlink():
                raise ValueError(f"Symlink in artifact: {path.relative_to(root)}")
            if path == root / ".git":
                dirs.remove(name)
        for name in names:
            path = Path(directory) / name
            if path == root / ".git":
                continue
            if path.is_symlink() or not path.is_file():
                raise ValueError(f"Non-regular artifact file: {path.relative_to(root)}")
            relative = relative_path(path.relative_to(root).as_posix())
            files[relative] = path.read_bytes()
    if not files:
        raise ValueError("Empty candidate artifact")
    return files


def compare(manifest, originals, candidate):
    validate_manifest(manifest, originals)
    if not candidate:
        raise ValueError("Empty candidate artifact")
    differences = []
    for entry in manifest["protected"]:
        path = entry["path"]
        if path not in candidate:
            differences.append({"path": path, "kind": "missing"})
        elif originals[path] != candidate[path]:
            before, after = originals[path], candidate[path]
            first = next((i for i, (a, b) in enumerate(zip(before, after)) if a != b), min(len(before), len(after)))
            differences.append({"path": path, "kind": "changed",
                                "baseline": metadata(path, before),
                                "candidate": metadata(path, after),
                                "firstDifferentByte": first,
                                "baselineHex": before[first:first + 32].hex(),
                                "candidateHex": after[first:first + 32].hex()})
    for path in SHARED:
        if path not in candidate:
            differences.append({"path": path, "kind": "missing shared integration file"})
    allowed = set(originals) | {e["path"] for e in manifest.get("allowedAdditions", [])}
    for path in sorted(set(candidate) - allowed):
        differences.append({"path": path, "kind": "unexpected addition"})
    return differences


def validate_provenance(provenance, manifest, input_root):
    if provenance.get("schemaVersion") != 1 or provenance.get("channel") != manifest["channel"]:
        raise ValueError("Provenance schema/channel differs from frozen baseline")
    if provenance.get("providerBaseRevision") != manifest["providerBaseRevision"]:
        raise ValueError("Provenance provider revision differs from frozen baseline")
    if not REVISION.fullmatch(provenance.get("generatorRevision", "")) or provenance.get("generatorWorkingTreeDirty") is not False:
        raise ValueError("Accepted provenance requires a clean, committed generator revision")
    if git(input_root, "rev-parse", "HEAD").decode().strip() != provenance["generatorRevision"]:
        raise ValueError("Input checkout differs from the recorded generator revision")
    if git(input_root, "status", "--porcelain", "--untracked-files=all").strip():
        raise ValueError("Generator input checkout is not clean")
    reproduction = provenance.get("baselineReproduction", {})
    if reproduction.get("status") != "reproduced":
        raise ValueError("Baseline reproduction remains blocked; no preservation acceptance claim is permitted")
    if reproduction.get("protectedFileCount") != len(manifest["protected"]):
        raise ValueError("Reproduction evidence has the wrong protected comparison count")
    for key in ("generatorBaseRevision", "stripeGoBaseRevision", "metronomeSourceRevision", "metronomeGoSourceRevision"):
        if not REVISION.fullmatch(provenance.get(key, "")):
            raise ValueError(f"Missing immutable provenance revision: {key}")
    for key in ("metronomeGoModuleVersion", "metronomeGoModuleChecksum"):
        if not provenance.get(key):
            raise ValueError(f"Missing SDK provenance: {key}")
    for name in ("node", "yarn", "go", "gofmt", "terraformPluginDocs"):
        if not provenance.get("toolchain", {}).get(name):
            raise ValueError(f"Missing toolchain version: {name}")
    inputs = provenance.get("inputs", {})
    for key in ("stripeV1", "stripeV2", "metronomeSpec", "metronomePolicies", "generatorLockfile"):
        item = inputs.get(key, {})
        path = relative_path(item.get("path"))
        if not item.get("revision") or not SHA256.fullmatch(item.get("sha256", "")):
            raise ValueError(f"Missing locked input provenance: {key}")
        actual = locked_input(input_root, path)
        if hashlib.sha256(actual).hexdigest() != item["sha256"]:
            raise ValueError(f"Locked input bytes differ: {key} ({path})")
    policies = provenance.get("allMetronomeInputs", [])
    if {item.get("path") for item in policies} != METRONOME_INPUTS or len(policies) != len(METRONOME_INPUTS):
        raise ValueError("Metronome input inventory is incomplete or duplicated")
    for item in policies:
        path = relative_path(item.get("path"))
        if hashlib.sha256(locked_input(input_root, path)).hexdigest() != item.get("sha256"):
            raise ValueError(f"Locked Metronome input differs: {path}")
    evidence = reproduction.get("evidence", {})
    path = relative_path(evidence.get("path"))
    actual = locked_input(input_root, path)
    if hashlib.sha256(actual).hexdigest() != evidence.get("sha256"):
        raise ValueError("Missing or changed baseline reproduction evidence")
    report = json.loads(actual)
    if (report.get("providerBaseRevision") != manifest["providerBaseRevision"]
            or report.get("protectedFileCount") != len(manifest["protected"])
            or report.get("changed") != [] or report.get("unexpectedAdditions") != []):
        raise ValueError("Baseline reproduction evidence does not prove exact output parity")
    for item in manifest["protected"]:
        if report.get("outputs", {}).get(item["path"]) != item["sha256"]:
            raise ValueError(f"Reproduction hash differs for {item['path']}")


def locked_input(root, relative):
    root = Path(root)
    if root.is_symlink() or not root.is_dir():
        raise ValueError("Locked input root must be a real directory")
    target = root
    for part in PurePosixPath(relative_path(relative)).parts:
        target = target / part
        if target.is_symlink():
            raise ValueError(f"Symlink in locked input path: {relative}")
    if not target.is_file():
        raise ValueError(f"Missing locked input: {relative}")
    return target.read_bytes()


def freeze(repo, revision, channel, additions):
    originals = git_files(repo, revision)
    if not set(SHARED) <= set(originals):
        raise ValueError("Baseline lacks required shared integration files")
    manifest = {
        "schemaVersion": 1, "channel": channel, "providerBaseRevision": revision,
        "protected": [metadata(p, originals[p]) for p in sorted(set(originals) - set(SHARED))],
        "sharedIntegrationFiles": [{**metadata(p, originals[p]), "purpose": purpose} for p, purpose in SHARED.items()],
        "allowedAdditions": additions,
    }
    validate_manifest(manifest, originals)
    return manifest


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    commands = parser.add_subparsers(dest="command", required=True)
    create = commands.add_parser("freeze", help="Create a manifest once; never overwrite it")
    create.add_argument("--baseline-repo", required=True)
    create.add_argument("--revision", required=True)
    create.add_argument("--channel", choices=sorted(CHANNELS), required=True)
    create.add_argument("--allowed-additions", help="Reviewed JSON array of exact path/owner/purpose entries")
    create.add_argument("--output", required=True)
    for command in ("diagnose", "check"):
        check = commands.add_parser(command, help="check requires proven baseline provenance; diagnose cannot establish acceptance")
        check.add_argument("--manifest", required=True)
        check.add_argument("--baseline-repo", required=True)
        candidate = check.add_mutually_exclusive_group(required=True)
        candidate.add_argument("--candidate")
        candidate.add_argument("--candidate-revision", help="Compare the exact committed tree in baseline-repo")
        if command == "check":
            check.add_argument("--provenance", required=True)
            check.add_argument("--input-root", required=True)
    args = parser.parse_args()
    try:
        if args.command == "freeze":
            additions = json.loads(Path(args.allowed_additions).read_text()) if args.allowed_additions else []
            manifest = freeze(args.baseline_repo, args.revision, args.channel, additions)
            with open(args.output, "x") as output:
                json.dump(manifest, output, indent=2)
                output.write("\n")
        else:
            manifest = json.loads(Path(args.manifest).read_text())
            originals = git_files(args.baseline_repo, manifest["providerBaseRevision"])
            candidate = git_files(args.baseline_repo, args.candidate_revision) if args.candidate_revision else artifact_files(args.candidate)
            differences = compare(manifest, originals, candidate)
            if args.command == "check":
                validate_provenance(json.loads(Path(args.provenance).read_text()), manifest, args.input_root)
            print(json.dumps({"mode": args.command, "acceptance": args.command == "check" and not differences,
                              "protectedFileCount": len(manifest["protected"]), "differences": differences}, indent=2))
            return 1 if differences else 0
    except (ValueError, KeyError, TypeError, OSError, subprocess.CalledProcessError) as error:
        print(f"Preservation gate failed: {error}", file=sys.stderr)
        return 2
    return 0


if __name__ == "__main__":
    sys.exit(main())

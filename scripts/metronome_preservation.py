#!/usr/bin/env python3
"""Compare provider artifacts directly with the base Git tree, without transforming bytes."""

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


ADDITION_DIRECTORIES = (
    "internal/provider/metronome/", "internal/metronomeclient/",
    "internal/providerdata/", "internal/combinedprovider/",
    "internal/terraformruntime/", "testdata/metronome/", "examples/metronome/",
)
ADDITION_FILES = {
    ".github/workflows/metronome-terraform.yml", ".metronome-generated-assets.json",
    "LICENSE.metronome", "LICENSE.metronome-sdk", "generation-provenance.json",
    "docs/guides/metronome.md", "scripts/metronome_offline.py",
    "scripts/metronome_preservation.py", "metronome_environment_test.go",
    "metronome_acceptance_test.go", "metronome_actions_acceptance_test.go",
    "metronome_customer_contract_acceptance_test.go", "metronome_migration_test.go",
    "metronome_bootstrap_test.go",
}


def allowed_addition(path):
    relative_path(path)
    return (path in ADDITION_FILES or path.startswith(ADDITION_DIRECTORIES)
            or re.fullmatch(r"docs/(resources|data-sources|actions)/metronome_[a-z0-9_]+\.md", path) is not None)


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


def compare(originals, candidate):
    if not set(SHARED) <= set(originals):
        raise ValueError("Base Git tree lacks required shared integration files")
    protected = set(originals) - set(SHARED)
    if not protected:
        raise ValueError("Empty protected comparison set")
    if not candidate:
        raise ValueError("Empty candidate artifact")
    for path in set(originals) | set(candidate):
        relative_path(path)
    differences = []
    for path in sorted(protected):
        if path not in candidate:
            differences.append({"path": path, "kind": "missing"})
        elif originals[path] != candidate[path]:
            before, after = originals[path], candidate[path]
            first = next((i for i, (a, b) in enumerate(zip(before, after)) if a != b), min(len(before), len(after)))
            differences.append({"path": path, "kind": "changed",
                                "baselineLength": len(before), "candidateLength": len(after),
                                "firstDifferentByte": first,
                                "baselineHex": before[first:first + 32].hex(),
                                "candidateHex": after[first:first + 32].hex()})
    for path in SHARED:
        if path not in candidate:
            differences.append({"path": path, "kind": "missing shared integration file"})
    for path in sorted(set(candidate) - set(originals)):
        if not allowed_addition(path):
            differences.append({"path": path, "kind": "unexpected addition"})
    return differences


def validate_provenance(provenance, base_revision, input_root):
    if provenance.get("schemaVersion") != 1 or provenance.get("channel") not in CHANNELS:
        raise ValueError("Invalid provenance schema/channel")
    if provenance.get("providerBaseRevision") != base_revision:
        raise ValueError("Provenance provider revision differs from the base Git commit")
    if not REVISION.fullmatch(provenance.get("generatorRevision", "")) or provenance.get("generatorWorkingTreeDirty") is not False:
        raise ValueError("Accepted provenance requires a clean, committed generator revision")
    if git(input_root, "rev-parse", "HEAD").decode().strip() != provenance["generatorRevision"]:
        raise ValueError("Input checkout differs from the recorded generator revision")
    if git(input_root, "status", "--porcelain", "--untracked-files=all").strip():
        raise ValueError("Generator input checkout is not clean")
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


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--baseline-repo", required=True)
    parser.add_argument("--base-revision", required=True, help="Full provider PR base commit")
    candidate = parser.add_mutually_exclusive_group(required=True)
    candidate.add_argument("--candidate", help="Final generated provider directory")
    candidate.add_argument("--candidate-revision", help="Exact candidate commit in baseline-repo")
    parser.add_argument("--provenance", help="Also verify generation inputs and clean generator identity")
    parser.add_argument("--input-root", help="Generator checkout used for provenance verification")
    args = parser.parse_args()
    if bool(args.provenance) != bool(args.input_root):
        parser.error("--provenance and --input-root must be supplied together")
    try:
        originals = git_files(args.baseline_repo, args.base_revision)
        candidate = git_files(args.baseline_repo, args.candidate_revision) if args.candidate_revision else artifact_files(args.candidate)
        differences = compare(originals, candidate)
        if args.provenance:
            recorded = Path(args.provenance).read_bytes()
            if recorded != candidate.get("generation-provenance.json"):
                raise ValueError("Provenance does not belong to the candidate artifact")
            validate_provenance(json.loads(recorded), args.base_revision, args.input_root)
        print(json.dumps({"baseRevision": args.base_revision,
                          "protectedFileCount": len(set(originals) - set(SHARED)),
                          "provenanceValidated": bool(args.provenance),
                          "differences": differences}, indent=2))
        return 1 if differences else 0
    except (ValueError, KeyError, TypeError, OSError, subprocess.CalledProcessError) as error:
        print(f"Stripe byte comparison failed: {error}", file=sys.stderr)
        return 2


if __name__ == "__main__":
    sys.exit(main())

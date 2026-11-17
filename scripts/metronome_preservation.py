#!/usr/bin/env python3
"""Compare provider artifacts directly with the base Git tree, without transforming bytes."""

import argparse
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
REVISION = re.compile(r"^[0-9a-f]{40}$")

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
    "internal/terraformruntime/", "examples/metronome/",
)
ADDITION_FILES = {
    ".github/workflows/metronome-terraform.yml", ".metronome-generated-assets.json",
    "LICENSE.metronome", "LICENSE.metronome-sdk",
    "docs/guides/metronome.md", "scripts/test_metronome.py",
    "scripts/metronome_preservation.py", "metronome_environment_test.go",
    "metronome_acceptance_test.go", "metronome_actions_acceptance_test.go",
    "metronome_customer_contract_acceptance_test.go",
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
    protected = {path for path in originals if path not in SHARED and not allowed_addition(path)}
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


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--baseline-repo", required=True)
    parser.add_argument("--base-revision", required=True, help="Full provider PR base commit")
    candidate = parser.add_mutually_exclusive_group(required=True)
    candidate.add_argument("--candidate", help="Final generated provider directory")
    candidate.add_argument("--candidate-revision", help="Exact candidate commit in baseline-repo")
    args = parser.parse_args()
    try:
        originals = git_files(args.baseline_repo, args.base_revision)
        candidate = git_files(args.baseline_repo, args.candidate_revision) if args.candidate_revision else artifact_files(args.candidate)
        differences = compare(originals, candidate)
        print(json.dumps({"baseRevision": args.base_revision,
                          "protectedFileCount": len([path for path in originals if path not in SHARED and not allowed_addition(path)]),
                          "differences": differences}, indent=2))
        return 1 if differences else 0
    except (ValueError, KeyError, TypeError, OSError, subprocess.CalledProcessError) as error:
        print(f"Stripe byte comparison failed: {error}", file=sys.stderr)
        return 2


if __name__ == "__main__":
    sys.exit(main())

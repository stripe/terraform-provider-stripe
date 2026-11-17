#!/usr/bin/env python3
"""Run the Terraform integration tests and reject missing or skipped suites."""

import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile


REQUIRED = {
    "TestMetronomeExplicitResourcesLifecycle",
    "TestMetronomeGeneratedReadAndExplicitActionInvocation",
    "TestMetronomeTerraformCustomerContractLifecycle",
    "TestMetronomeBridgeModuleOrderingAndTargetedBootstrap",
    "TestMetronomeInstalledProviderWithoutCredentials",
}


def main():
    root = Path(__file__).resolve().parents[1]
    environment = {
        name: value for name, value in os.environ.items()
        if not name.startswith(("TF_", "STRIPE_", "METRONOME_"))
    }
    environment["METRONOME_REQUIRE_ACTION_TEST"] = "1"
    version = json.loads(subprocess.check_output(
        ["terraform", "version", "-json"], env=environment, cwd=root,
    ))["terraform_version"]
    if version != "1.16.5":
        raise ValueError(f"Required Terraform version is 1.16.5, got {version}")
    with tempfile.TemporaryDirectory(prefix="metronome-tests-") as directory:
        log = Path(directory) / "tests.jsonl"
        with log.open("wb") as output:
            result = subprocess.run(
                ["go", "test", "-json", "-count=1", "-run", "^TestMetronome", "-timeout", "30m", "./internal/acctest/cases/metronome"],
                cwd=root, env=environment, stdout=output, stderr=subprocess.PIPE,
            )
        events = [json.loads(line) for line in log.read_text().splitlines() if line]
        passed = {e["Test"] for e in events if e.get("Action") == "pass" and "Test" in e}
        skipped = sorted({e["Test"] for e in events if e.get("Action") == "skip" and e.get("Test", "").startswith("TestMetronome")})
        missing = sorted(REQUIRED - passed)
        if result.returncode or missing or skipped:
            print(result.stderr.decode(), file=sys.stderr)
            for event in events:
                if event.get("Output"):
                    print(event["Output"], end="", file=sys.stderr)
            raise ValueError(f"Terraform tests failed: exit={result.returncode}, missing={missing}, skipped={skipped}")
        report = {
            "status": "passed", "terraform": version,
            "requiredSuites": sorted(REQUIRED), "passedTestsAndSubtests": len(passed),
            "skippedRequiredTests": skipped,
            "liveAPIs": False,
        }
        print(json.dumps(report, indent=2))
        report_path = os.environ.get("METRONOME_TEST_REPORT")
        if report_path:
            Path(report_path).write_text(json.dumps(report, indent=2) + "\n")


if __name__ == "__main__":
    try:
        main()
    except (ValueError, OSError, subprocess.CalledProcessError) as error:
        print(f"Metronome tests failed: {error}", file=sys.stderr)
        sys.exit(1)

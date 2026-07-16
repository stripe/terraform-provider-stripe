// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const managedModuleConfigFileName = "main.tf"

type managedModuleCleanupCandidate struct {
	config string
	env    []string
}

var terraformIsolatedEnvKeys = []string{
	"TF_CLI_CONFIG_FILE",
	"TF_DATA_DIR",
	"TF_PLUGIN_CACHE_DIR",
	"TF_REATTACH_PROVIDERS",
}

var (
	currentProviderBinaryBuildOnce sync.Once
	currentProviderBinaryPath      string
	currentProviderBinaryErr       error
)

func RunManagedModuleLegacyUpgradeCase(
	t *testing.T,
	tc ManagedCase,
	legacyProviderVersion string,
) {
	t.Helper()

	if !AcceptanceEnabled() {
		t.Skip("set TF_ACC=1 to run Terraform acceptance tests")
	}
	if !ShouldRunCase(tc.Definition) {
		t.Skip("case filtered out by selectors")
	}

	env := ResolveTestEnv(t, tc.Definition.Group, tc.Definition.RequiredEnv)
	ApplyResolvedAccountEnv(t, env)
	client := NewStripeClient(env)
	terraformPath := terraformCLIPath(t)
	caseReplacements := map[string]string{
		"{{RAND}}": fmt.Sprintf("%d", time.Now().UnixNano()),
	}
	createConfig := RenderManagedConfig(env, tc.CreateTemplate, caseReplacements)
	legacyConfig := legacyStripeRequiredProvidersConfig(legacyProviderVersion) + createConfig
	currentConfig := currentStripeRequiredProvidersConfig() + createConfig
	workdir := t.TempDir()
	legacyTerraformRCPath := filepath.Join(workdir, "legacy.terraformrc")
	if err := writeDirectTerraformCLIConfig(legacyTerraformRCPath); err != nil {
		t.Fatal(err)
	}
	legacyEnv := terraformCommandEnv(
		map[string]string{
			"TF_CLI_CONFIG_FILE": legacyTerraformRCPath,
		},
		terraformIsolatedEnvKeys...,
	)
	cleanupCandidates := []managedModuleCleanupCandidate{
		{
			config: legacyConfig,
			env:    legacyEnv,
		},
	}
	needsCleanup := false

	t.Cleanup(func() {
		if !needsCleanup {
			return
		}

		cleanupErrors := []string{}
		for _, candidate := range cleanupCandidates {
			if err := writeManagedModuleConfig(workdir, candidate.config); err != nil {
				cleanupErrors = append(
					cleanupErrors,
					fmt.Sprintf("write config failed: %v", err),
				)
				continue
			}
			if err := terraformInitReconfigure(terraformPath, workdir, candidate.env); err != nil {
				cleanupErrors = append(
					cleanupErrors,
					fmt.Sprintf("terraform init failed: %v", err),
				)
				continue
			}
			if err := terraformDestroyAutoApprove(terraformPath, workdir, candidate.env); err != nil {
				cleanupErrors = append(
					cleanupErrors,
					fmt.Sprintf("terraform destroy failed: %v", err),
				)
				continue
			}

			return
		}

		t.Errorf(
			"module legacy-upgrade cleanup failed:\n%s",
			strings.Join(cleanupErrors, "\n"),
		)
	})

	if err := writeManagedModuleConfig(workdir, legacyConfig); err != nil {
		t.Fatal(err)
	}
	if err := terraformInitReconfigure(terraformPath, workdir, legacyEnv); err != nil {
		t.Fatal(err)
	}
	if err := terraformApplyAutoApprove(terraformPath, workdir, legacyEnv); err != nil {
		t.Fatal(err)
	}
	needsCleanup = true

	providerBinaryPath := currentProviderBinaryForModuleUpgrade(t)
	currentTerraformRCPath := filepath.Join(workdir, "current.terraformrc")
	if err := writeDevOverrideTerraformCLIConfig(
		currentTerraformRCPath,
		filepath.Dir(providerBinaryPath),
	); err != nil {
		t.Fatal(err)
	}
	currentEnv := terraformCommandEnv(
		map[string]string{
			"TF_CLI_CONFIG_FILE": currentTerraformRCPath,
		},
		terraformIsolatedEnvKeys...,
	)
	cleanupCandidates = append(
		[]managedModuleCleanupCandidate{
			{
				config: currentConfig,
				env:    currentEnv,
			},
		},
		cleanupCandidates...,
	)

	if err := writeManagedModuleConfig(workdir, currentConfig); err != nil {
		t.Fatal(err)
	}
	if err := terraformInitReconfigure(terraformPath, workdir, currentEnv); err != nil {
		t.Fatal(err)
	}
	if err := terraformPlanMustBeNoop(terraformPath, workdir, currentEnv); err != nil {
		t.Fatal(err)
	}
	if err := terraformApplyAutoApprove(terraformPath, workdir, currentEnv); err != nil {
		t.Fatal(err)
	}
	if err := terraformPlanMustBeNoop(terraformPath, workdir, currentEnv); err != nil {
		t.Fatal(err)
	}

	state, err := terraformStateFromShowJSON(terraformPath, workdir, currentEnv)
	if err != nil {
		t.Fatal(err)
	}
	if err := verifyManagedCaseStateAddresses(state, tc); err != nil {
		t.Fatal(err)
	}
	if verifyCreate := stateVerifierCheck(env, client, tc.VerifyCreate); verifyCreate != nil {
		if err := verifyCreate(state); err != nil {
			t.Fatal(err)
		}
	}

	if err := terraformDestroyAutoApprove(terraformPath, workdir, currentEnv); err != nil {
		t.Fatal(err)
	}
	needsCleanup = false

	if verifyDestroy := stateVerifierCheck(env, client, tc.VerifyDestroy); verifyDestroy != nil {
		if err := verifyDestroy(state); err != nil {
			t.Fatal(err)
		}
	}
}

func terraformCLIPath(t *testing.T) string {
	t.Helper()

	if terraformPath := os.Getenv("TF_ACC_TERRAFORM_PATH"); terraformPath != "" {
		return terraformPath
	}

	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		t.Fatalf("failed to locate terraform binary: %v", err)
	}

	return terraformPath
}

func currentProviderBinaryForModuleUpgrade(t *testing.T) string {
	t.Helper()

	currentProviderBinaryBuildOnce.Do(func() {
		buildDir, err := os.MkdirTemp("", "stripe-provider-module-upgrade-*")
		if err != nil {
			currentProviderBinaryErr = fmt.Errorf(
				"failed to create provider build dir: %w",
				err,
			)
			return
		}

		currentProviderBinaryPath = filepath.Join(
			buildDir,
			terraformProviderBinaryName(),
		)
		output, exitCode, err := runCommand(
			moduleUpgradeProviderRepoRoot(),
			terraformCommandEnv(nil),
			"go",
			"build",
			"-o",
			currentProviderBinaryPath,
			".",
		)
		if err != nil {
			currentProviderBinaryErr = commandFailureError(
				"go build",
				exitCode,
				output,
				err,
			)
		}
	})

	if currentProviderBinaryErr != nil {
		t.Fatal(currentProviderBinaryErr)
	}

	return currentProviderBinaryPath
}

func terraformProviderBinaryName() string {
	if runtime.GOOS == "windows" {
		return "terraform-provider-stripe.exe"
	}

	return "terraform-provider-stripe"
}

func moduleUpgradeProviderRepoRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to resolve provider repo root")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", "..", ".."))
}

func writeManagedModuleConfig(workdir string, config string) error {
	configPath := filepath.Join(workdir, managedModuleConfigFileName)
	if err := os.WriteFile(configPath, []byte(config), 0o600); err != nil {
		return fmt.Errorf("failed to write module config %s: %w", configPath, err)
	}

	return nil
}

func writeDirectTerraformCLIConfig(terraformRCPath string) error {
	terraformRC := `provider_installation {
  direct {}
}
`
	if err := os.WriteFile(terraformRCPath, []byte(terraformRC), 0o600); err != nil {
		return fmt.Errorf("failed to write terraform rc %s: %w", terraformRCPath, err)
	}

	return nil
}

func writeDevOverrideTerraformCLIConfig(
	terraformRCPath string,
	providerBinaryDir string,
) error {
	terraformRC := fmt.Sprintf(
		`provider_installation {
  dev_overrides {
    "registry.terraform.io/stripe/stripe" = %q
  }
  direct {}
}
`,
		providerBinaryDir,
	)
	if err := os.WriteFile(terraformRCPath, []byte(terraformRC), 0o600); err != nil {
		return fmt.Errorf("failed to write terraform rc %s: %w", terraformRCPath, err)
	}

	return nil
}

func terraformCommandEnv(overrides map[string]string, removals ...string) []string {
	removalSet := map[string]struct{}{}
	for _, key := range removals {
		removalSet[key] = struct{}{}
	}

	env := []string{}
	for _, entry := range os.Environ() {
		key, _, found := strings.Cut(entry, "=")
		if found {
			if _, ok := removalSet[key]; ok {
				continue
			}
		}
		env = append(env, entry)
	}

	for key, value := range overrides {
		replaced := false
		for index, entry := range env {
			if strings.HasPrefix(entry, key+"=") {
				env[index] = key + "=" + value
				replaced = true
				break
			}
		}
		if !replaced {
			env = append(env, key+"="+value)
		}
	}

	return env
}

func terraformInitReconfigure(terraformPath string, workdir string, env []string) error {
	output, exitCode, err := runCommand(
		workdir,
		env,
		terraformPath,
		"init",
		"-reconfigure",
	)
	if err != nil {
		return commandFailureError("terraform init -reconfigure", exitCode, output, err)
	}

	return nil
}

func terraformApplyAutoApprove(terraformPath string, workdir string, env []string) error {
	output, exitCode, err := runCommand(
		workdir,
		env,
		terraformPath,
		"apply",
		"-auto-approve",
	)
	if err != nil {
		return commandFailureError("terraform apply -auto-approve", exitCode, output, err)
	}

	return nil
}

func terraformDestroyAutoApprove(terraformPath string, workdir string, env []string) error {
	output, exitCode, err := runCommand(
		workdir,
		env,
		terraformPath,
		"destroy",
		"-auto-approve",
	)
	if err != nil {
		return commandFailureError("terraform destroy -auto-approve", exitCode, output, err)
	}

	return nil
}

func terraformPlanMustBeNoop(terraformPath string, workdir string, env []string) error {
	output, exitCode, err := runCommand(
		workdir,
		env,
		terraformPath,
		"plan",
		"-detailed-exitcode",
	)
	if err == nil {
		return nil
	}
	if exitCode == 2 {
		trimmedOutput := strings.TrimSpace(output)
		if trimmedOutput == "" {
			return fmt.Errorf("terraform plan detected drift or pending changes")
		}

		return fmt.Errorf(
			"terraform plan detected drift or pending changes:\n%s",
			trimmedOutput,
		)
	}

	return commandFailureError("terraform plan -detailed-exitcode", exitCode, output, err)
}

func terraformStateFromShowJSON(
	terraformPath string,
	workdir string,
	env []string,
) (*terraform.State, error) {
	output, exitCode, err := runCommand(
		workdir,
		env,
		terraformPath,
		"show",
		"-json",
	)
	if err != nil {
		return nil, commandFailureError("terraform show -json", exitCode, output, err)
	}

	jsonState := &tfjson.State{}
	jsonState.UseJSONNumber(true)
	if err := json.Unmarshal([]byte(output), jsonState); err != nil {
		return nil, fmt.Errorf(
			"failed to decode terraform show -json output: %w\n%s",
			err,
			strings.TrimSpace(output),
		)
	}

	return terraformStateFromJSON(jsonState)
}

func verifyManagedCaseStateAddresses(state *terraform.State, tc ManagedCase) error {
	for _, address := range managedCaseStateAddresses(tc) {
		if _, err := ResourcePrimaryID(state, address); err != nil {
			return err
		}
	}

	return nil
}

func runCommand(
	workdir string,
	env []string,
	command string,
	args ...string,
) (string, int, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = workdir
	cmd.Env = env
	output, err := cmd.CombinedOutput()
	if err == nil {
		return string(output), 0, nil
	}

	exitCode := -1
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		exitCode = exitErr.ExitCode()
	}

	return string(output), exitCode, err
}

func commandFailureError(
	command string,
	exitCode int,
	output string,
	err error,
) error {
	trimmedOutput := strings.TrimSpace(output)
	if exitCode >= 0 {
		if trimmedOutput == "" {
			return fmt.Errorf("%s failed with exit code %d: %w", command, exitCode, err)
		}

		return fmt.Errorf(
			"%s failed with exit code %d: %w\n%s",
			command,
			exitCode,
			err,
			trimmedOutput,
		)
	}
	if trimmedOutput == "" {
		return fmt.Errorf("%s failed: %w", command, err)
	}

	return fmt.Errorf("%s failed: %w\n%s", command, err, trimmedOutput)
}

func terraformStateFromJSON(jsonState *tfjson.State) (*terraform.State, error) {
	// The existing verifier helpers operate on terraform.State, so the module
	// runner shims Terraform CLI JSON back into that legacy structure.
	state := terraform.NewState() //nolint:staticcheck
	state.TFVersion = jsonState.TerraformVersion

	if jsonState.Values == nil {
		return state, nil
	}

	rootModule := state.RootModule()
	if rootModule.Resources == nil {
		rootModule.Resources = map[string]*terraform.ResourceState{}
	}
	if rootModule.Outputs == nil {
		rootModule.Outputs = map[string]*terraform.OutputState{}
	}
	for key, output := range jsonState.Values.Outputs {
		outputState, err := terraformOutputStateFromJSON(output)
		if err != nil {
			return nil, err
		}
		rootModule.Outputs[key] = outputState
	}
	if err := appendTerraformModuleResources(rootModule, jsonState.Values.RootModule); err != nil {
		return nil, err
	}

	return state, nil
}

func appendTerraformModuleResources(
	rootModule *terraform.ModuleState,
	module *tfjson.StateModule,
) error {
	if module == nil {
		return nil
	}

	for _, resource := range module.Resources {
		resourceState, err := terraformResourceStateFromJSON(resource)
		if err != nil {
			return err
		}
		rootModule.Resources[resource.Address] = resourceState
	}
	for _, childModule := range module.ChildModules {
		if err := appendTerraformModuleResources(rootModule, childModule); err != nil {
			return err
		}
	}

	return nil
}

func terraformOutputStateFromJSON(output *tfjson.StateOutput) (*terraform.OutputState, error) {
	outputState := &terraform.OutputState{
		Sensitive: output.Sensitive,
	}

	switch value := output.Value.(type) {
	case string:
		outputState.Type = "string"
		outputState.Value = value
		return outputState, nil
	case []interface{}:
		outputState.Type = "list"
		if len(value) == 0 {
			outputState.Value = value
			return outputState, nil
		}

		switch firstElement := value[0].(type) {
		case string:
			elements := make([]interface{}, len(value))
			for index, element := range value {
				stringElement, ok := element.(string)
				if !ok {
					outputState.Value = value
					return outputState, nil
				}
				elements[index] = stringElement
			}
			outputState.Value = elements
		case bool:
			elements := make([]interface{}, len(value))
			for index, element := range value {
				boolElement, ok := element.(bool)
				if !ok {
					outputState.Value = value
					return outputState, nil
				}
				elements[index] = boolElement
			}
			outputState.Value = elements
		case json.Number:
			elements := make([]interface{}, len(value))
			for index, element := range value {
				numberElement, ok := element.(json.Number)
				if !ok {
					outputState.Value = value
					return outputState, nil
				}
				elements[index] = numberElement
			}
			outputState.Value = elements
		case []interface{}:
			outputState.Value = value
		case map[string]interface{}:
			outputState.Value = value
		default:
			return nil, fmt.Errorf(
				"unexpected output list element type: %T",
				firstElement,
			)
		}

		return outputState, nil
	case map[string]interface{}:
		outputState.Type = "map"
		outputState.Value = value
		return outputState, nil
	case bool:
		outputState.Type = "string"
		outputState.Value = strconv.FormatBool(value)
		return outputState, nil
	case json.Number:
		outputState.Type = "string"
		outputState.Value = value.String()
		return outputState, nil
	}

	return nil, fmt.Errorf("unexpected output type: %T", output.Value)
}

func terraformResourceStateFromJSON(
	resource *tfjson.StateResource,
) (*terraform.ResourceState, error) {
	flatmap := &terraformStateFlatmap{}
	if err := flatmap.FromMap(resource.AttributeValues); err != nil {
		return nil, err
	}
	attributes := flatmap.Flatmap()

	instanceStateID, ok := attributes["id"]
	if !ok {
		instanceStateID = "id-attribute-not-set"
	}

	return &terraform.ResourceState{
		Provider: resource.ProviderName,
		Type:     resource.Type,
		Primary: &terraform.InstanceState{
			ID:         instanceStateID,
			Attributes: attributes,
			Meta: map[string]interface{}{
				"schema_version": int(resource.SchemaVersion),
			},
			Tainted: resource.Tainted,
		},
		Dependencies: resource.DependsOn,
	}, nil
}

type terraformStateFlatmap struct {
	values map[string]string
}

func (flatmap *terraformStateFlatmap) FromMap(attributes map[string]interface{}) error {
	if flatmap.values == nil {
		flatmap.values = make(map[string]string, len(attributes))
	}

	return flatmap.AddMap("", attributes)
}

func (flatmap *terraformStateFlatmap) AddMap(
	prefix string,
	attributes map[string]interface{},
) error {
	for key, value := range attributes {
		flatKey := key
		if prefix != "" {
			flatKey = fmt.Sprintf("%s.%s", prefix, key)
		}

		if err := flatmap.AddEntry(flatKey, value); err != nil {
			return fmt.Errorf("unable to add map key %q entry: %w", flatKey, err)
		}
	}

	mapLengthKey := "%"
	if prefix != "" {
		mapLengthKey = fmt.Sprintf("%s.%s", prefix, "%")
	}

	if err := flatmap.AddEntry(mapLengthKey, strconv.Itoa(len(attributes))); err != nil {
		return fmt.Errorf(
			"unable to add map length %q entry: %w",
			mapLengthKey,
			err,
		)
	}

	return nil
}

func (flatmap *terraformStateFlatmap) AddSlice(
	name string,
	elements []interface{},
) error {
	for index, element := range elements {
		key := fmt.Sprintf("%s.%d", name, index)
		if err := flatmap.AddEntry(key, element); err != nil {
			return fmt.Errorf("unable to add slice key %q entry: %w", key, err)
		}
	}

	sliceLengthKey := fmt.Sprintf("%s.#", name)
	if err := flatmap.AddEntry(sliceLengthKey, strconv.Itoa(len(elements))); err != nil {
		return fmt.Errorf(
			"unable to add slice length %q entry: %w",
			sliceLengthKey,
			err,
		)
	}

	return nil
}

func (flatmap *terraformStateFlatmap) AddEntry(
	key string,
	value interface{},
) error {
	switch element := value.(type) {
	case nil:
		return nil
	case bool:
		flatmap.values[key] = strconv.FormatBool(element)
	case json.Number:
		flatmap.values[key] = element.String()
	case string:
		flatmap.values[key] = element
	case map[string]interface{}:
		if err := flatmap.AddMap(key, element); err != nil {
			return err
		}
	case []interface{}:
		if err := flatmap.AddSlice(key, element); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%q: unexpected type (%T)", key, element)
	}

	return nil
}

func (flatmap *terraformStateFlatmap) Flatmap() map[string]string {
	return flatmap.values
}

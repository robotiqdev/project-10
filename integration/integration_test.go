//go:build integration

package integration_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var binaryPath string

func TestMain(m *testing.M) {
	// Build the binary into a temp file so tests can invoke it.
	dir, err := os.MkdirTemp("", "calc-integration-*")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(dir)

	binaryPath = filepath.Join(dir, "calc_test_bin")

	// Build from the module root (one level up from integration/).
	build := exec.Command("go", "build", "-o", binaryPath, ".")
	build.Dir = filepath.Join("..")
	out, buildErr := build.CombinedOutput()
	if buildErr != nil {
		panic("failed to build binary: " + buildErr.Error() + "\n" + string(out))
	}

	os.Exit(m.Run())
}

// runCalc executes the compiled binary with the provided arguments and returns
// stdout, stderr, and the exit code.
func runCalc(args ...string) (stdout, stderr string, exitCode int) {
	cmd := exec.Command(binaryPath, args...)
	var outBuf, errBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	stdout = outBuf.String()
	stderr = errBuf.String()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = -1
		}
	}
	return
}

// TestCLI_Addition verifies that "3 + 4" prints "3 + 4 = 7" to stdout and exits 0.
func TestCLI_Addition(t *testing.T) {
	stdout, _, exitCode := runCalc("3", "+", "4")

	if exitCode != 0 {
		t.Errorf("exit code = %d, want 0", exitCode)
	}
	want := "3 + 4 = 7"
	if !strings.Contains(stdout, want) {
		t.Errorf("stdout = %q, want it to contain %q", stdout, want)
	}
}

// TestCLI_DivisionByZero verifies that "10 / 0" writes a "division by zero"
// message to stderr and exits 1.
func TestCLI_DivisionByZero(t *testing.T) {
	_, stderr, exitCode := runCalc("10", "/", "0")

	if exitCode != 1 {
		t.Errorf("exit code = %d, want 1", exitCode)
	}
	if !strings.Contains(strings.ToLower(stderr), "division by zero") {
		t.Errorf("stderr = %q, want it to contain 'division by zero'", stderr)
	}
}

// TestCLI_InvalidOperand verifies that a non-numeric operand ("abc + 1") causes
// the binary to exit with code 1.
func TestCLI_InvalidOperand(t *testing.T) {
	_, _, exitCode := runCalc("abc", "+", "1")

	if exitCode != 1 {
		t.Errorf("exit code = %d, want 1 for invalid operand", exitCode)
	}
}

// TestCLI_MissingArgs verifies that invoking the binary with no arguments
// causes it to exit with code 1.
func TestCLI_MissingArgs(t *testing.T) {
	_, _, exitCode := runCalc()

	if exitCode != 1 {
		t.Errorf("exit code = %d, want 1 for missing arguments", exitCode)
	}
}

// TestCLI_TooFewArgs verifies that too few arguments (e.g. just one token)
// also cause the binary to exit with code 1.
func TestCLI_TooFewArgs(t *testing.T) {
	_, _, exitCode := runCalc("3")

	if exitCode != 1 {
		t.Errorf("exit code = %d, want 1 for too few arguments", exitCode)
	}
}

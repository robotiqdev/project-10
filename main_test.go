// Package main_test provides integration tests for the calculator binary entry
// point (main.go). These tests exercise the full wiring layer: argument parsing,
// calculation, and output/error routing as described in TASK-4511.
//
// NOTE: Deep integration tests (e.g. full golden-file coverage) are handled in
// TASK-4514. This file covers the smoke-test layer: confirming that the binary
// exits correctly and routes output to stdout/stderr as required by Unix
// conventions.
package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// binaryPath is the path to the compiled calculator binary, built once in TestMain.
var binaryPath string

// TestMain builds the binary once before running any tests, so every test case
// can invoke the real executable rather than go run (which is slower and
// unsuitable for subprocess-based integration tests).
func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "calc-integration-*")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(dir)

	binaryPath = filepath.Join(dir, "calculator")
	if runtime.GOOS == "windows" {
		binaryPath += ".exe"
	}

	// Determine the module root (the directory that contains main.go).
	_, thisFile, _, _ := runtime.Caller(0)
	moduleRoot := filepath.Dir(thisFile)

	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	cmd.Dir = moduleRoot
	if out, err := cmd.CombinedOutput(); err != nil {
		panic("failed to build binary: " + string(out))
	}

	os.Exit(m.Run())
}

// runCalc is a test helper that invokes the calculator binary with the supplied
// arguments and returns stdout, stderr, and whether the process exited cleanly.
func runCalc(t *testing.T, args ...string) (stdout, stderr string, exitOK bool) {
	t.Helper()
	var outBuf, errBuf bytes.Buffer
	cmd := exec.Command(binaryPath, args...)
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	return strings.TrimRight(outBuf.String(), "\n\r"),
		strings.TrimRight(errBuf.String(), "\n\r"),
		err == nil
}

// --- Happy-path tests ---

// TestBinary_Addition verifies that "3 + 4" exits 0 and writes the result to stdout.
func TestBinary_Addition(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "3", "+", "4")
	if !ok {
		t.Fatalf("3 + 4: expected exit 0, got non-zero; stderr=%q", stderr)
	}
	if stdout == "" {
		t.Error("3 + 4: expected non-empty stdout, got empty")
	}
	if !strings.Contains(stdout, "7") {
		t.Errorf("3 + 4: expected stdout to contain %q, got %q", "7", stdout)
	}
	if stderr != "" {
		t.Errorf("3 + 4: expected empty stderr, got %q", stderr)
	}
}

// TestBinary_Subtraction verifies that "10 - 3" exits 0 and writes the result to stdout.
func TestBinary_Subtraction(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "10", "-", "3")
	if !ok {
		t.Fatalf("10 - 3: expected exit 0, got non-zero; stderr=%q", stderr)
	}
	if !strings.Contains(stdout, "7") {
		t.Errorf("10 - 3: expected stdout to contain %q, got %q", "7", stdout)
	}
	if stderr != "" {
		t.Errorf("10 - 3: expected empty stderr, got %q", stderr)
	}
}

// TestBinary_Multiplication verifies that "3 * 4" exits 0 and writes the result to stdout.
func TestBinary_Multiplication(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "3", "*", "4")
	if !ok {
		t.Fatalf("3 * 4: expected exit 0, got non-zero; stderr=%q", stderr)
	}
	if !strings.Contains(stdout, "12") {
		t.Errorf("3 * 4: expected stdout to contain %q, got %q", "12", stdout)
	}
	if stderr != "" {
		t.Errorf("3 * 4: expected empty stderr, got %q", stderr)
	}
}

// TestBinary_Division verifies that "10 / 2" exits 0 and writes the result to stdout.
func TestBinary_Division(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "10", "/", "2")
	if !ok {
		t.Fatalf("10 / 2: expected exit 0, got non-zero; stderr=%q", stderr)
	}
	if !strings.Contains(stdout, "5") {
		t.Errorf("10 / 2: expected stdout to contain %q, got %q", "5", stdout)
	}
	if stderr != "" {
		t.Errorf("10 / 2: expected empty stderr, got %q", stderr)
	}
}

// TestBinary_StdoutNotEmpty_AllOperators runs a table-driven check that every
// supported operator yields non-empty stdout and empty stderr on valid input.
func TestBinary_StdoutNotEmpty_AllOperators(t *testing.T) {
	cases := []struct {
		left, op, right string
	}{
		{"5", "+", "3"},
		{"10", "-", "4"},
		{"6", "*", "7"},
		{"15", "/", "3"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.left+" "+tc.op+" "+tc.right, func(t *testing.T) {
			stdout, stderr, ok := runCalc(t, tc.left, tc.op, tc.right)
			if !ok {
				t.Fatalf("expected exit 0; stderr=%q", stderr)
			}
			if stdout == "" {
				t.Error("expected non-empty stdout")
			}
			if stderr != "" {
				t.Errorf("expected empty stderr, got %q", stderr)
			}
		})
	}
}

// --- Error-path tests ---

// TestBinary_DivisionByZero verifies that "5 / 0" exits non-zero and writes
// the error to stderr, not stdout. This is the primary manually-verified case
// documented in TASK-4511.
func TestBinary_DivisionByZero(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "5", "/", "0")
	if ok {
		t.Fatal("5 / 0: expected non-zero exit, got exit 0")
	}
	if stderr == "" {
		t.Error("5 / 0: expected error message on stderr, got empty")
	}
	// The error must mention "division by zero" (case-insensitive) so the user
	// receives a meaningful message.
	if !strings.Contains(strings.ToLower(stderr), "division by zero") {
		t.Errorf("5 / 0: stderr %q does not contain %q", stderr, "division by zero")
	}
	if stdout != "" {
		t.Errorf("5 / 0: expected empty stdout on error, got %q", stdout)
	}
}

// TestBinary_NoArguments verifies that omitting all arguments exits non-zero
// and writes an error to stderr.
func TestBinary_NoArguments(t *testing.T) {
	stdout, stderr, ok := runCalc(t)
	if ok {
		t.Fatal("no args: expected non-zero exit, got exit 0")
	}
	if stderr == "" {
		t.Error("no args: expected error on stderr, got empty")
	}
	if stdout != "" {
		t.Errorf("no args: expected empty stdout, got %q", stdout)
	}
}

// TestBinary_TooFewArguments verifies that fewer than three arguments causes
// a non-zero exit with an error on stderr.
func TestBinary_TooFewArguments(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "3", "+")
	if ok {
		t.Fatal("too few args: expected non-zero exit, got exit 0")
	}
	if stderr == "" {
		t.Error("too few args: expected error on stderr, got empty")
	}
	if stdout != "" {
		t.Errorf("too few args: expected empty stdout, got %q", stdout)
	}
}

// TestBinary_TooManyArguments verifies that more than three arguments causes
// a non-zero exit with an error on stderr.
func TestBinary_TooManyArguments(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "3", "+", "4", "5")
	if ok {
		t.Fatal("too many args: expected non-zero exit, got exit 0")
	}
	if stderr == "" {
		t.Error("too many args: expected error on stderr, got empty")
	}
	if stdout != "" {
		t.Errorf("too many args: expected empty stdout, got %q", stdout)
	}
}

// TestBinary_InvalidLeftNumber verifies that a non-numeric left argument causes
// a non-zero exit with an error on stderr.
func TestBinary_InvalidLeftNumber(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "abc", "+", "4")
	if ok {
		t.Fatal("invalid left number: expected non-zero exit, got exit 0")
	}
	if stderr == "" {
		t.Error("invalid left number: expected error on stderr, got empty")
	}
	if stdout != "" {
		t.Errorf("invalid left number: expected empty stdout, got %q", stdout)
	}
}

// TestBinary_InvalidRightNumber verifies that a non-numeric right argument causes
// a non-zero exit with an error on stderr.
func TestBinary_InvalidRightNumber(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "3", "+", "xyz")
	if ok {
		t.Fatal("invalid right number: expected non-zero exit, got exit 0")
	}
	if stderr == "" {
		t.Error("invalid right number: expected error on stderr, got empty")
	}
	if stdout != "" {
		t.Errorf("invalid right number: expected empty stdout, got %q", stdout)
	}
}

// TestBinary_UnknownOperator verifies that an unsupported operator causes
// a non-zero exit with an error on stderr.
func TestBinary_UnknownOperator(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "3", "%", "4")
	if ok {
		t.Fatal("unknown operator: expected non-zero exit, got exit 0")
	}
	if stderr == "" {
		t.Error("unknown operator: expected error on stderr, got empty")
	}
	if stdout != "" {
		t.Errorf("unknown operator: expected empty stdout, got %q", stdout)
	}
}

// TestBinary_ErrorGoesToStderrNotStdout verifies the Unix convention that error
// messages are directed to stderr and never leak to stdout.
func TestBinary_ErrorGoesToStderrNotStdout(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{"no args", []string{}},
		{"division by zero", []string{"5", "/", "0"}},
		{"invalid number", []string{"abc", "+", "1"}},
		{"unknown operator", []string{"1", "^", "2"}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, ok := runCalc(t, tc.args...)
			if ok {
				t.Fatal("expected non-zero exit")
			}
			if stdout != "" {
				t.Errorf("stdout must be empty on error, got %q", stdout)
			}
			if stderr == "" {
				t.Error("stderr must contain error message")
			}
		})
	}
}

// TestBinary_SuccessGoesToStdoutNotStderr verifies that successful calculations
// write only to stdout and leave stderr empty.
func TestBinary_SuccessGoesToStdoutNotStderr(t *testing.T) {
	stdout, stderr, ok := runCalc(t, "3", "+", "4")
	if !ok {
		t.Fatalf("expected exit 0; stderr=%q", stderr)
	}
	if stdout == "" {
		t.Error("expected result on stdout, got empty")
	}
	if stderr != "" {
		t.Errorf("expected empty stderr on success, got %q", stderr)
	}
}

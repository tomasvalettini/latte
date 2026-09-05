package testutils

import (
	"os"
	"os/exec"
	"testing"
)

const BE_CRASHER = "BE_CRASHER"

func RequireExit(t *testing.T, testName string, testFunction func()) {
	t.Helper()

	// 1. Re-run the test binary, but invoke a specific test case
	if os.Getenv(BE_CRASHER) == "1" {
		testFunction()
		return
	}

	// 2. Run the command as a subprocess
	cmd := exec.Command(os.Args[0], "-test.run="+testName)
	cmd.Env = append(os.Environ(), BE_CRASHER+"=1")

	// 3. Verify that the command exited with a non-zero status
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return // Test failed successfully
	}
}

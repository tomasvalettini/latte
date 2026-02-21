package testutils

import (
	"log"
	"testing"
)

func exitIfError() {
	log.Fatalf("Fatal error occurred")
}

func TestRequireExit(t *testing.T) {
	RequireExit(t, "TestRequireExit", exitIfError)
}

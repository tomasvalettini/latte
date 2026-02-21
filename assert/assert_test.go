package assert

import (
	"testing"

	testutils "github.com/tomasvalettini/latte/test-utils"
)

func TestAssertTrue(t *testing.T) {
	Assert(true, "no message needed")
}

func TestAssertFalseWithoutPanic(t *testing.T) {
	testutils.RequireExit("TestAssertFalseWithoutPanic", assertFalseTesting)
}

func assertFalseTesting() {
	Assert(false, "Oh no!!!")
}

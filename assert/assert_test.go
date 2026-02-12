package assert

import (
	"testing"
)

func TestAssertTrue(t *testing.T) {
	Assert(true, "no message needed")
}

func TestAssertFalse(t *testing.T) {
	errorMessage := "Oh no!!!"

	defer func() {
		if r := recover(); r != nil {
			if r == errorMessage {
				t.Logf("Recovered from panic \n")
			} else {
				t.Errorf("Not the panic we were expecting")
			}
		}
	}()

	Assert(false, errorMessage)
}


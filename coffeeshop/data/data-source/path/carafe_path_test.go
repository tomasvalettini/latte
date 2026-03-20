package carafepath

import (
	"strings"
	"testing"

	"github.com/tomasvalettini/latte/assert"
)

func TestGettingCarafePath(t *testing.T) {
	performTest(&LocalCarafePath{}, CARAFE_FILE_NAME, LATTE_HOME_DIRECTORY)
}

func TestGettingTestCarafePath(t *testing.T) {
	performTest(GetTestingCarafePath(), "test.json", "latte")
}

func performTest(carafePath CarafePath, expectedFile string, expectedDirectory string) {
	pathSections := strings.Split(carafePath.GetCarafePath(), "/")
	size := len(pathSections)

	assert.Assert(pathSections[size-1] == expectedFile, "Carafe file name is not what was expected!")
	assert.Assert(pathSections[size-2] == expectedDirectory, "Latte directory is not what was expected!")
}

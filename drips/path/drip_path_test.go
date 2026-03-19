package drippath

import (
	"strings"
	"testing"

	"github.com/tomasvalettini/latte/assert"
)

func TestGettingDripPath(t *testing.T) {
	performTest(&LocalDripPath{}, DRIP_FILE_NAME, LATTE_HOME_DIRECTORY)
}

func TestGettingTestDripPath(t *testing.T) {
	performTest(GetTestingDripPath(), "test.json", "latte")
}

func performTest(dripPath DripPath, expectedFile string, expectedDirectory string) {
	pathSections := strings.Split(dripPath.GetDripPath(), "/")
	size := len(pathSections)

	assert.Assert(pathSections[size-1] == expectedFile, "Drip file name is not what was expected!")
	assert.Assert(pathSections[size-2] == expectedDirectory, "Latte directory is not what was expected!")
}

package taskpath

import (
	"strings"
	"testing"

	"github.com/tomasvalettini/latte/assert"
)

func TestGettingTaskPath(t *testing.T) {
	performTest(&LocalTaskPath{}, TASK_FILE_NAME, LATTE_HOME_DIRECTORY)
}

func TestGettingTestTaskPath(t *testing.T) {
	performTest(GetTestingTaskPath(), "test.json", "latte")
}

func performTest(taskPath TaskPath, expectedFile string, expectedDirectory string) {
	pathSections := strings.Split(taskPath.GetTaskPath(), "/")
	size := len(pathSections)

	assert.Assert(pathSections[size-1] == expectedFile, "Task file name is not what was expected!")
	assert.Assert(pathSections[size-2] == expectedDirectory, "Latte directory is not what was expected!")
}

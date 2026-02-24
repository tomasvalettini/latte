package backlog

import (
	"strings"
	"testing"

	"github.com/tomasvalettini/latte/assert"
)

func TestGettingTaskPath(t *testing.T) {
	localTaskPath := LocalTaskPath{}
	pathSections := strings.Split(localTaskPath.GetTaskPath(), "/")
	size := len(pathSections)

	assert.Assert(pathSections[size-1] == TASK_FILE_NAME, "Task file name is not what was expected!")
	assert.Assert(pathSections[size-2] == LATTE_HOME_DIRECTORY, "Latte directory is not what was expected!")
}

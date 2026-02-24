package backlog

import (
	"os"
	"path/filepath"
)

const LATTE_HOME_DIRECTORY = ".latte"
const TASK_FILE_NAME = "tasks.json"

type LocalTaskPath struct{}

func (ltp *LocalTaskPath) GetTaskPath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, LATTE_HOME_DIRECTORY, TASK_FILE_NAME)
}

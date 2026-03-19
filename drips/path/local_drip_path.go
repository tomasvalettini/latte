package drippath

import (
	"os"
	"path/filepath"
)

const LATTE_HOME_DIRECTORY = ".latte"
const DRIP_FILE_NAME = "drips.json"

type LocalDripPath struct{}

func (ltp *LocalDripPath) GetDripPath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, LATTE_HOME_DIRECTORY, DRIP_FILE_NAME)
}

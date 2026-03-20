package carafepath

import (
	"os"
	"path/filepath"
)

const LATTE_HOME_DIRECTORY = ".latte"
const CARAFE_FILE_NAME = "carafes.json"

type LocalCarafePath struct{}

func (ltp *LocalCarafePath) GetCarafePath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, LATTE_HOME_DIRECTORY, CARAFE_FILE_NAME)
}

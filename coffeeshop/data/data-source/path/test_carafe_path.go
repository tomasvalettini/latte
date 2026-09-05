package carafepath

const TMP = "tmp/"
const TEST_CARAFE_FILE = "latte/test.json"

type TestCarafePath struct{}

func (ttp *TestCarafePath) GetCarafePath() string {
	return TMP + TEST_CARAFE_FILE
}

func GetTestingCarafePath() CarafePath {
	return &TestCarafePath{}
}

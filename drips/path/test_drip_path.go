package drippath

const TMP = "tmp/"
const TEST_DRIP_FILE = "latte/test.json"

type TestDripPath struct{}

func (ttp *TestDripPath) GetDripPath() string {
	return TMP + TEST_DRIP_FILE
}

func GetTestingDripPath() DripPath {
	return &TestDripPath{}
}

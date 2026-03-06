package taskpath

const TMP = "tmp/"
const TEST_TASK_FILE = "latte/test.json"

type TestTaskPath struct{}

func (ttp *TestTaskPath) GetTaskPath() string {
	return TMP + TEST_TASK_FILE
}

func GetTestingTaskPath() TaskPath {
	return &TestTaskPath{}
}

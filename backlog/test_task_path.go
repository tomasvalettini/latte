package backlog

const TMP = "tmp/"

type TestTaskPath struct{}

func (ttp *TestTaskPath) GetTaskPath() string {
	return TMP + "latte/test.json"
}

func GetTestingTaskPath() TaskPath {
	return &TestTaskPath{}
}

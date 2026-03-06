package taskdatasource

import taskdatamodel "github.com/tomasvalettini/latte/tasks/data/model"

type TaskDataSource interface {
	Load() []taskdatamodel.Task
	Save(tasks []taskdatamodel.Task)
}

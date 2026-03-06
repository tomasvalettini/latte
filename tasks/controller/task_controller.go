package controller

import (
	"fmt"
	"log"
	"strconv"

	taskdatasource "github.com/tomasvalettini/latte/tasks/data/data-source"
	taskdatamodel "github.com/tomasvalettini/latte/tasks/data/model"
	taskpath "github.com/tomasvalettini/latte/tasks/path"
)

type TaskController struct {
	taskPath   taskpath.TaskPath
	dataSource *taskdatasource.TaskBacklog
}

func NewTaskController(taskPath taskpath.TaskPath) *TaskController {
	dataSource := taskdatasource.NewTaskBacklog(taskPath.GetTaskPath())

	return &TaskController{
		taskPath:   taskPath,
		dataSource: dataSource,
	}
}

func (tc *TaskController) ListTasks() {
	tasks, notified := tc.loadTasks()
	if !notified {
		return
	}

	fmt.Println("===========")
	fmt.Println(" TASK LIST ")
	fmt.Println("===========")

	w := taskdatamodel.MaxIdWidth(tasks)
	for _, t := range tasks {
		fmt.Printf("  [%*d]  %s\n", w, t.Id, t.Text)
	}
}

func (tc *TaskController) AddTask(taskText string) {
	tasks := tc.dataSource.Load()
	nextId := taskdatamodel.GetNextId(tasks)

	task := taskdatamodel.Task{
		Id:   nextId,
		Text: taskText,
	}
	tasks = append(tasks, task)

	tc.dataSource.Save(tasks)
	fmt.Println("Task added successfully!!!")
}

func (tc *TaskController) DeleteTask(idStr string) {
	id, ok := tc.getTaskId(idStr)
	if !ok {
		return
	}

	tasks, notified := tc.loadTasks()
	if !notified {
		return
	}

	idx := taskdatamodel.FindIndexFromId(tasks, id)
	tasks = append(tasks[:idx], tasks[idx+1:]...)
	tc.dataSource.Save(tasks)

	fmt.Printf("Task with id: %d was successfully removed!!\n", id)
}

func (tc *TaskController) UpdateTask(idStr string, newText string) {
	id, ok := tc.getTaskId(idStr)
	if !ok {
		return
	}

	tasks := tc.dataSource.Load()
	index := taskdatamodel.FindIndexFromId(tasks, id)

	tasks[index].Text = newText
	tc.dataSource.Save(tasks)

	fmt.Printf("Task with id: %d was successfully modified with new text:%s\n", id, newText)
}

func (tc *TaskController) loadTasks() ([]taskdatamodel.Task, bool) {
	tasks := tc.dataSource.Load()
	tasksCount := len(tasks)

	if tasksCount <= 0 {
		fmt.Println("No tasks yet.")
		return nil, false
	}

	return tasks, true
}

func (tc *TaskController) getTaskId(idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalln("Id entered is not a number.")
	}

	return id, true
}

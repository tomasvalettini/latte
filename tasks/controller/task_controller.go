package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tomasvalettini/latte/backlog"
)

type TaskController struct {
	taskPath   backlog.TaskPath
	dataSource *backlog.Backlog
}

func NewTaskController(taskPath backlog.TaskPath) *TaskController {
	dataSource := backlog.NewBacklog(taskPath.GetTaskPath())

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

	w := backlog.MaxIdWidth(tasks)
	for _, t := range tasks {
		fmt.Printf("  [%*d]  %s\n", w, t.Id, t.Text)
	}
}

func (tc *TaskController) AddTask(taskText string) {
	tasks := tc.dataSource.Load()
	nextId := backlog.GetNextId(tasks)

	task := backlog.Task{
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

	idx := backlog.FindIndexFromId(tasks, id)
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
	index := backlog.FindIndexFromId(tasks, id)

	tasks[index].Text = newText
	tc.dataSource.Save(tasks)

	fmt.Printf("Task with id: %d was successfully modified with new text:%s\n", id, newText)
}

func (tc *TaskController) loadTasks() ([]backlog.Task, bool) {
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

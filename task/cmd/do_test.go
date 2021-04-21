package cmd

import (
	"errors"
	"testing"

	"github.ibm.com/gophercises/task/db"
)

func TestDoCmd (t *testing.T) {
	tempList := allTasksHandler
	tempDelete := deletetaskHandler
	defer func() {
		allTasksHandler = tempList
		deletetaskHandler = tempDelete
	}()
	allTasksHandler = func() ([]db.Task, error) {
		var tasks []db.Task
		tasks = append(tasks, db.Task{Key: 1, Value: "completed task"})
		return tasks, nil
	}
	deletetaskHandler = func(key int) error {
		return nil
	}
	doCmd.Run(doCmd, []string{"1"})
}

func TestDoCmdFail(t *testing.T) {
	tempList := allTasksHandler
	defer func() {
		allTasksHandler = tempList
	}()
	allTasksHandler = func() ([]db.Task, error) {
		var tasks []db.Task
		tasks = append(tasks, db.Task{Key: 1, Value: "Completed task"})
		return tasks, errors.New("Error")
	}
	doCmd.Run(doCmd, []string{"1"})
}

func TestDoCmdWithChar(t *testing.T) {
	tempList := allTasksHandler
	defer func() {
		allTasksHandler = tempList
	}()
	allTasksHandler = func() ([]db.Task, error) {
		var tasks []db.Task
		tasks = append(tasks, db.Task{Key: 1, Value: "Completed task"})
		return tasks, errors.New("Error")
	}
	doCmd.Run(doCmd, []string{"A"})
}

func TestFailInDelete(t *testing.T) {
	tempList := allTasksHandler
	tempDelete := deletetaskHandler
	defer func() {
		allTasksHandler = tempList
		deletetaskHandler = tempDelete
	}()
	allTasksHandler = func() ([]db.Task, error) {
		var tasks []db.Task
		tasks = append(tasks, db.Task{Key: 1, Value: "completed task"})
		return tasks, nil
	}
	deletetaskHandler = func(key int) error {
		return errors.New("Error")
	}

	doCmd.Run(doCmd, []string{"1", "2"})

}



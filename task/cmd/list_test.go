package cmd

import (
	"github.ibm.com/gophercises/task/db"
	"errors"
	"testing"
 )

func TestListCmd(t *testing.T) {
	temp:= allTasksHandler
	defer func (){
		allTasksHandler = temp
	}()
	allTasksHandler = func ()([]db.Task, error){
		var t []db.Task
		t = append(t, db.Task{Key:1 ,Value: "Complete the exercise."})
		return t, nil
	}
	listCmd.Run(listCmd, []string{})
}
func TestListcmdNegative(t *testing.T){
	temp := allTasksHandler
	defer func (){
		allTasksHandler = temp
	}()
	allTasksHandler =func ()([]db.Task, error){
		var t []db.Task
		t = append(t, db.Task{Key:1 ,Value: "Complete the exercise."})
		return t, errors.New("Error")
	}
	listCmd.Run(listCmd,[] string{})
}
func TestWithZeroTask(t *testing.T) {
	tempList := allTasksHandler
	defer func (){
		allTasksHandler = tempList
	}()	
	allTasksHandler = func ()([]db.Task, error) {
		var tasks []db.Task
		return tasks, nil
	}
	listCmd.Run(listCmd, []string{})
}

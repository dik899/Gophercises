package cmd

import (
	"testing"
	"errors"
)


func TestAddcmd(t *testing.T) {
	temp := createtaskHandler
	defer func (){
		createtaskHandler = temp
	}()
	createtaskHandler= func (task string)(int, error){
		return 1,nil
	}
	addCmd.Run(addCmd, []string{"Complete the exercise"})
  
}
func TestAddcmdNegative(t *testing.T) {
	temp := createtaskHandler 
	defer func() {
		createtaskHandler = temp
	}()
	createtaskHandler = func(task string)(int, error){
		return 1, errors.New("Error")
	}
	addCmd.Run(addCmd, []string{"complete the exercise"})
  
}



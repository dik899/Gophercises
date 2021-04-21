package cmd

import(

	"testing"
     "errors"
	
)

func TestGetCmd(t *testing.T) { 
	defer func(){
		
	}()
	var1 = func(task string)(string, error){
		return "pass", nil
	}
	getCmd.Run(getCmd, []string{"124"})
	
}
func TestGetCmdFail(t *testing.T){
	defer func(){
		
	}()
	var1 = func(task string)(string, error){
		return "pass", errors.New("Error")
	}
	getCmd.Run(getCmd,[]string{"124"})
}
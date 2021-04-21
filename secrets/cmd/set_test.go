package cmd

import(
	"testing"
	"errors"
	
)
func TestSetCmd(t *testing.T) { 
	
	setCmd.Run(setCmd, []string{"facebook","12345"})
	
	
}
func TestSetCmdFail(t *testing.T){
	
	defer func(){
		
	}()
	sethandler = func(key string, value string)( error){
		return errors.New("error")
	}
	setCmd.Run(setCmd, []string{"facebook","12345"})
}
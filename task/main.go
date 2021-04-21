package main

import (
	//"fmt"
	"path/filepath"

	"github.ibm.com/gophercises/task/cmd"
	"github.ibm.com/gophercises/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {

	startApp()
}

func startApp() error {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		return err
	}
	return cmd.RootCmd.Execute()
}

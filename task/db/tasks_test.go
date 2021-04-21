package db

import (
	"os"
	"path/filepath"
	"testing"
     "fmt"
	"github.com/mitchellh/go-homedir"
)

func TestMain(m *testing.M) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	x := m.Run()
	err := os.Remove(dbPath)
	if err != nil {
		return
	}
	os.Exit(x)
}
func TestInit(t *testing.T) {
	resultError := Init(loadFile())
	if resultError != nil {
		fmt.Println("Errror")
	}
	db.Close()
}

func TestInitFail(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home,"/invalid/task.db")
	err := Init(dbPath)
	if err != nil {
		fmt.Println("Error")
	}
}

func TestCreateTask(t *testing.T) {
	Init(loadFile())

	_, err := CreateTask("Gophercises")
    if err != nil {
		fmt.Println("Error")
	}
	db.Close()
}

func TestAllTasks(t *testing.T) {
	Init(loadFile())
	_, err := AllTasks()
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("PASS")
	}
	db.Close()
}

func TestDeleteTask(t *testing.T) {
	Init(loadFile())
	err := DeleteTask(1)
	if err != nil {
		fmt.Println("Error")
	}
	db.Close()
}

func loadFile() string {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	return dbPath
}

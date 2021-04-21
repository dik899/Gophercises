package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.ibm.com/gophercises/task/db"
)

var (
	createtaskHandler = db.CreateTask
)
var addCmd = &cobra.Command{
	//command name
	Use: "add",
	// short description of command
	Short: "Add a task to your task list.",
	// function involked while add command execution
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		_, err := createtaskHandler(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

// func createTask(task string) (int, error) {
// 	return db.CreateTask(task)
// }

//init() can run with or before main()
// bootstrap the addcommand
func init() {
	RootCmd.AddCommand(addCmd)
}

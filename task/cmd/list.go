package cmd
import (
	"fmt"
	//"os"

	"github.ibm.com/gophercises/task/db"
	"github.com/spf13/cobra"
)
var(
allTasksHandler =db.AllTasks
)
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := allTasksHandler()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			//os.Exit(1)
			return 
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!Why not take a vacation?")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}
// func allTasks() ([]db.Task, error){
// 	return db.AllTasks()
// }

func init() {
	RootCmd.AddCommand(listCmd)
}

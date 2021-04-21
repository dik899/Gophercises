package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	secret "github.ibm.com/gophercises/secret/vault"
)

var (
	sethandler = secret.File(encodingKey, secretsPath()).Set
)
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {

		key, value := args[0], args[1]
		err := sethandler(key, value)
		if err != nil {
			fmt.Println("error")
		}

		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}

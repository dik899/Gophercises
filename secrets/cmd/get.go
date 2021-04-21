package cmd

import (
	"fmt"
    "github.ibm.com/gophercises/secret/vault"
	"github.com/spf13/cobra"
)
var(
	var1 = secret.File(encodingKey, secretsPath()).Get
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		
		key := args[0]
		value, err := var1(key)
		if err != nil {
			fmt.Println("no value set")
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}


func init() {
	RootCmd.AddCommand(getCmd)
}

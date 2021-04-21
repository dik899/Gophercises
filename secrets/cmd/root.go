package cmd

import (
	"path/filepath"
    //"fmt"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
)
//RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is an API key and other secrets manager",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "the key to use when encoding and decoding secrets")
  // RootCmd.PersistentFlags().StringVarP(&encodingKey, "value", "V", "", " The value of key")

}


func secretsPath() string {
	home, _ := homedir.Dir()
	//fmt.Println("Home directory is ::",home)
	return filepath.Join(home, ".secrets")
}

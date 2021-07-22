package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"ls"},
	Short: "installed version list",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(cfg.GorootsDir)
		cobra.CheckErr(err)

		for _, file := range files {
			if file.IsDir() == false {
				continue
			}
			fmt.Println(file.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

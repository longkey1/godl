package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gorootsCmd represents the goroots command
var gorootsCmd = &cobra.Command{
	Use:   "goroots",
	Short: "describe goroots directory path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfg.GorootsDir)
	},
}

func init() {
	rootCmd.AddCommand(gorootsCmd)
}

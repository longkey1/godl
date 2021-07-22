package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove specific version",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("requires a version argument.")
		}
		ver := args[0]
		versionDir := filepath.Join(cfg.GorootsDir, ver)
		err := os.RemoveAll(versionDir)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

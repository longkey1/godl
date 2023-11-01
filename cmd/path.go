package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "describe path",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(cfg.GorootsDir)
			return
		}
		target := args[0]

		latest := latestVersion(target, localLatestVersions())
		if latest == InitialVersion {
			log.Fatalln("Not found matched version")
		}

		fmt.Println(filepath.Join(cfg.GorootsDir, latest))
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}

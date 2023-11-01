package cmd

import (
	"fmt"
	hv "github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
)

var listRemoteCmd = &cobra.Command{
	Use:     "list-remote",
	Aliases: []string{"ls-remote"},
	Short:   "downloadable version list",
	Run: func(cmd *cobra.Command, args []string) {
		latest, err := cmd.Flags().GetBool("latest")
		cobra.CheckErr(err)

		var versions []*hv.Version
		if latest {
			versions = remoteLatestVersions()
		} else {
			versions = remoteVersions()
		}

		// After this, the versions are properly sorted
		for _, v := range versions {
			fmt.Println(v.Original())
		}
	},
}

func init() {
	rootCmd.AddCommand(listRemoteCmd)
	listRemoteCmd.Flags().Bool("latest", false, "filter latest version only")
}

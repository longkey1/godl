package cmd

import (
	"fmt"
	hv "github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
	"io/ioutil"
	"sort"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "installed version list",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(cfg.GorootsDir)
		cobra.CheckErr(err)

		var versionsRaw []string
		for _, file := range files {
			if file.IsDir() == false {
				continue
			}
			versionsRaw = append(versionsRaw, file.Name())
		}

		versions := make([]*hv.Version, len(versionsRaw))
		for i, raw := range versionsRaw {
			v, _ := hv.NewVersion(raw)
			versions[i] = v
		}

		sort.Sort(sort.Reverse(hv.Collection(versions)))
		for _, v := range versions {
			fmt.Println(v.Original())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

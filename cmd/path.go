package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/spf13/cobra"
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

		files, err := ioutil.ReadDir(cfg.GorootsDir)
		cobra.CheckErr(err)

		latest := ""
		for _, file := range files {
			if file.IsDir() == false {
				continue
			}
			if strings.Index(file.Name(), target) != 0 {
				continue
			}
			if file.Name() == target {
				latest = file.Name()
				break
			}
			targetVer, err := semver.Make(file.Name())
			if err != nil {
				continue
			}
			if len(latest) == 0 {
				latest = targetVer.String()
			}
			latestVer, _ := semver.Make(latest)
			if latestVer.Compare(targetVer) < 0 {
				latest = targetVer.String()
			}
		}
		if len(latest) == 0 {
			log.Fatalln("Not found matched version")
		}

		fmt.Println(filepath.Join(cfg.GorootsDir, latest))
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}

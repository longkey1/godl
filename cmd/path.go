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

		latestVer := ""
		for _, file := range files {
			if file.IsDir() == false {
				continue
			}
			if strings.Index(file.Name(), target) != 0 {
				continue
			}
			if len(latestVer) == 0 {
				latestVer = file.Name()
			}
			ver1, err1 := semver.Make(latestVer)
			ver2, err2 := semver.Make(file.Name())
			if err1 != nil || err2 != nil {
				continue
			}
			if ver1.Compare(ver2) < 0 {
				latestVer = file.Name()
			}
		}
		if len(latestVer) == 0 {
			log.Fatalln("Not found matched version")
		}

		fmt.Println(filepath.Join(cfg.GorootsDir, latestVer))
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}

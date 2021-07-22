/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("requires a version argument.")
		}
		ver := args[0]
		ext := "tar.gz"
		if runtime.GOARCH == "windows" {
			ext = "zip"
		}
		url := fmt.Sprintf("%s/dl/go%s.%s-%s.%s", GolangUrl, ver, runtime.GOOS, runtime.GOARCH, ext)

		res, err := http.Get(url)
		if err != nil {
			log.Fatalf("download error. %s", err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatalf("file close error. %s", err)
			}
		}(res.Body)

		archiveFile := filepath.Join(os.TempDir(), fmt.Sprintf("godl-archive.%s", ext))
		archive, err := os.Create(archiveFile)
		if err != nil {
			log.Fatalf("temporary file error. %s", err)
		}
		defer func(out *os.File) {
			err := out.Close()
			if err != nil {
				log.Fatalf("file close error. %s", err)
			}
		}(archive)

		_, err = io.Copy(archive, res.Body)

		extractDir := filepath.Join(GodlDir, fmt.Sprintf("godl-extract-%s", ver))
		err = archiver.Unarchive(archiveFile, extractDir)
		if err != nil {
			log.Fatalf("unarchive error. %s", err)
		}

		err = os.Rename(filepath.Join(extractDir, "go"), filepath.Join(GodlDir, ver))
		if err != nil {
			log.Fatalf("rename error. from: %s, to: %s, %s", filepath.Join(extractDir, "go"), filepath.Join(GodlDir, ver), err)
		}

		err = os.Remove(archiveFile)
		if err != nil {
			log.Fatalf("remove error. %s, %s", archiveFile, err)
		}

		err = os.Remove(extractDir)
		if err != nil {
			log.Fatalf("remove error. %s, %s", archiveFile, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

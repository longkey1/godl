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
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// listRemoteCmd represents the listRemote command
var listRemoteCmd = &cobra.Command{
	Use:   "list-remote",
	Aliases: []string{"ls-remote"},
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get(GolangUrl+"/dl")
		if err != nil {
			log.Fatal(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		versionMap := map[string]string{}
		doc.Find("a.download").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			url, _ := s.Attr("href")
			if strings.HasSuffix(url, "src.tar.gz") == false {
				return
			}
			reg := regexp.MustCompile(`/dl/go([0-9.]+)\.src\.tar\.gz$`)
			ver := reg.FindStringSubmatch(url)
			//fmt.Printf("%#v\n", ver)
			if len(ver) > 1 {
				versionMap[ver[1]] = ver[1]
			}
		})

		versions := []string{}
		for v := range versionMap {
			versions = append(versions, v)
		}

		sort.Strings(versions)
		for _, v := range versions {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listRemoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listRemoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listRemoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

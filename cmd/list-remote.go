package cmd

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var listRemoteCmd = &cobra.Command{
	Use:     "list-remote",
	Aliases: []string{"ls-remote"},
	Short:   "downloadable version list",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get(cfg.GolangUrl + "/dl")
		cobra.CheckErr(err)

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			cobra.CheckErr(err)
		}(res.Body)

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		cobra.CheckErr(err)

		versionMap := map[string]string{}
		doc.Find("a.download").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			url, _ := s.Attr("href")
			if strings.HasSuffix(url, "src.tar.gz") == false {
				return
			}
			reg := regexp.MustCompile(`/dl/go([0-9.]+)\.src\.tar\.gz$`)
			ver := reg.FindStringSubmatch(url)
			if len(ver) > 1 {
				versionMap[ver[1]] = ver[1]
			}
		})

		var versions []string
		for v := range versionMap {
			versions = append(versions, v)
		}

		latest, err := cmd.Flags().GetBool("latest")
		cobra.CheckErr(err)

		if latest {
			latestVersionMap := make(map[string]string)
			for v := range versionMap {
				fullVersion := v
				if (strings.Count(fullVersion, ".")) == 1 {
					fullVersion += ".0"
				}
				fullVersionSeparated := strings.Split(fullVersion, ".")
				minorVersion := strings.Join([]string{fullVersionSeparated[0], fullVersionSeparated[1]}, ".")
				_, has := latestVersionMap[minorVersion]
				if has == false {
					latestVersionMap[minorVersion] = fullVersion
				} else {
					for k, v := range latestVersionMap {
						if k != minorVersion {
							continue
						}
						latestVersionSeparated := strings.Split(v, ".")
						fullVersionPatchNumber, _ := strconv.Atoi(fullVersionSeparated[2])
						latestVersionPatchNumber, _ := strconv.Atoi(latestVersionSeparated[2])
						if fullVersionPatchNumber > latestVersionPatchNumber {
							latestVersionMap[minorVersion] = fullVersion
						}
					}
				}
			}
			var latestVersions []string
			for _, v := range latestVersionMap {
				latestVersions = append(latestVersions, v)
			}
			versions = latestVersions
		}

		sort.Strings(versions)
		for _, v := range versions {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listRemoteCmd)
	listRemoteCmd.Flags().Bool("latest", false, "filter latest version only")
}

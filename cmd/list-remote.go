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
	Short: "downloadable version list",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get(cfg.GolangUrl+"/dl")
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
}

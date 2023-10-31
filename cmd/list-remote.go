package cmd

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	hv "github.com/hashicorp/go-version"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
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

		var versionsRaw []string
		doc.Find("a.download").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			url, _ := s.Attr("href")
			if strings.HasSuffix(url, "src.tar.gz") == false {
				return
			}
			reg := regexp.MustCompile(`/dl/go([0-9.]+)\.src\.tar\.gz$`)
			ver := reg.FindStringSubmatch(url)
			if len(ver) > 1 {
				versionsRaw = append(versionsRaw, ver[1])
			}
		})

		versions := make([]*hv.Version, len(versionsRaw))
		for i, raw := range versionsRaw {
			v, _ := hv.NewVersion(raw)
			versions[i] = v
		}

		// After this, the versions are properly sorted
		sort.Sort(sort.Reverse(hv.Collection(versions)))

		latest, err := cmd.Flags().GetBool("latest")
		cobra.CheckErr(err)

		if latest {
			latestVersionsMap := make(map[string]*hv.Version)
			for _, v := range versions {
				seg := v.Segments()
				minorVersionKey := fmt.Sprintf("%d_%d", seg[0], seg[1])
				if _, ok := latestVersionsMap[minorVersionKey]; ok {
					continue
				}
				latestVersionsMap[minorVersionKey] = v
			}
			var latestVersions []*hv.Version
			for _, v := range latestVersionsMap {
				latestVersions = append(latestVersions, v)
			}
			versions = latestVersions
		}

		sort.Sort(sort.Reverse(hv.Collection(versions)))
		for _, v := range versions {
			fmt.Println(v.Original())
		}
	},
}

func init() {
	rootCmd.AddCommand(listRemoteCmd)
	listRemoteCmd.Flags().Bool("latest", false, "filter latest version only")
}

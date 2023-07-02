package cmd

import (
	"fmt"
	scrapper "github/Yarlaw07/Courls/pkg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "courls",
	Short: "Counter of urls in web domain",
	Long:  "Scraper counter of all urls in this web domain",
	Run: func(cmd *cobra.Command, args []string) {

		url := getUrl(args)

		ex, err := os.Executable()

		if err != nil {
			panic(err)
		}

		fmt.Println(filepath.Dir(ex))

		file, _ := os.Create("file.txt")

		c := scrapper.GetScrapper(url, file)

		fmt.Fprintln(file)

		defer file.Close()

		c.Visit(url)
	},
}

func getUrl(args []string) string {

	if len(args) != 1 {
		log.Fatalln("courls must accept only one parameter that a link to site")
	}

	url := args[0]

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		log.Fatalln("courls must have before link a http or https")
	}

	return url
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("filepath", "f", "", "specify filepath to resFile")
}

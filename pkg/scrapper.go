package scrapper

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetScrapper(url string, file *os.File) *colly.Collector {

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		href := e.Attr("href")

		if strings.HasPrefix(href, url) {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {

		fmt.Fprintln(file, r.URL.String())
	})

	return c
}

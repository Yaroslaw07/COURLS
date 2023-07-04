package scrapper

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

var Count = 0

func GetScrapper(url string, file *os.File, limit int) *colly.Collector {

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		
		fmt.Println(Count, limit, Count < limit)

		if Count >= limit {
			return
		}

		href := (e.Attr("href"))

		if strings.HasPrefix(href, url) {

			e.Request.Visit(clearUrl(href))
		}
	})

	c.OnRequest(func(r *colly.Request) {

		fmt.Fprintln(file, r.URL.String())
		Count++
	})

	return c
}

// clearUrl clear url parameters
func clearUrl(url string) string {

	lastIndex := strings.LastIndex(url, "/")

	return url[:lastIndex+1]
}

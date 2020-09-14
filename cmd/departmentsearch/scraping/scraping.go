package scraping

import (
	"braiton/braiton-home/cmd/internal/scraping/buscadorprop"

	"github.com/gocolly/colly/v2"
)

func MakeScrapingUrl(url string) {
	c := colly.NewCollector()
	_ = buscadorprop.MakeScraping(c, url)
	// fmt.Println(departments)
}

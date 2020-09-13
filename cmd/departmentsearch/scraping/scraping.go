package scraping

import (
	"braiton/braiton-home/cmd/internal/scraping/buscadorprop"
	"fmt"

	"github.com/gocolly/colly/v2"
)

func MakeScrapingUrl(url string) {
	c := colly.NewCollector()
	departments := buscadorprop.MakeScraping(c, url)
	fmt.Println(departments)
}

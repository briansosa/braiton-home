package buscadorprop

import (
	"braiton/braiton-home/cmd/internal/scraping"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func MakeScraping(c *colly.Collector, url string) []scraping.Department {
	var departments []scraping.Department
	c.OnHTML(".publicacion-item ", func(e *colly.HTMLElement) {
		department := scraping.Department{
			Address:      GetAddress(e),
			Locality:     GetLocality(e),
			Title:        GetTitle(e),
			Details:      GetDetails(e),
			Price:        GetPrice(e),
			LinkToDetail: GetLinkToDetail(e),
			Image:        GetImageName(e),
		}
		departments = append(departments, department)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)
	return departments
}

func GetAddress(e *colly.HTMLElement) string {
	return e.ChildText(".calle")
}

func GetLocality(e *colly.HTMLElement) string {
	return e.ChildText(".localidad")
}

func GetTitle(e *colly.HTMLElement) string {
	return e.ChildText(".content h2")
}

func GetDetails(e *colly.HTMLElement) string {
	return e.ChildText(".descripcion p")
}

func GetPrice(e *colly.HTMLElement) string {
	return e.ChildText(".precio")
}

func GetLinkToDetail(e *colly.HTMLElement) string {
	link := e.ChildAttr(".col-content > a", "href")
	absoluteLink := e.Request.AbsoluteURL(link)
	return absoluteLink
}

func GetImageLink(e *colly.HTMLElement) string {
	return e.ChildAttr(".col-image img", "src")
}

func GetImageName(e *colly.HTMLElement) string {
	linkToImage := GetImageLink(e)
	splitedLink := strings.Split(linkToImage, "/")
	imageName := splitedLink[len(splitedLink)-1]
	return imageName
}

package departmentsearch

import (
	"braiton/braiton-home/cmd/departmentsearch/routes"
	"braiton/braiton-home/cmd/departmentsearch/scraping"
)

func Run() {
	urlToScraping := "https://www.buscadorprop.com.ar/departamentos-loft-semipiso-alquiler-lanus-este-lanus-oeste-hasta-20000-pesos-mas-nuevas"
	scraping.MakeScrapingUrl(urlToScraping)
	routes.ConfigRoutes()
}

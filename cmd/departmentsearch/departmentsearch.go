package departmentsearch

import (
	"braiton/braiton-home/cmd/departmentsearch/routes"
	"braiton/braiton-home/cmd/departmentsearch/scraping"
	"braiton/braiton-home/cmd/internal/platform/postgres"
)

func Run() {
	// config db
	db, err := postgres.ConfigDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// get urls to be consulted

	// make scrapping
	urlToScraping := "https://www.buscadorprop.com.ar/departamentos-loft-semipiso-alquiler-lanus-este-lanus-oeste-hasta-20000-pesos-mas-nuevas"
	scraping.MakeScrapingUrl(urlToScraping)
	// compareIfNotExisting
	// manage files
	// put on db
	// config routes
	routes.ConfigRoutes()

}

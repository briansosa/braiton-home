package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type Departamento struct {
	Direccion    string
	Inmobiliaria string
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("api/departamentos-consultados", DepartamentosConsultados)

	router.Run(":" + port)
}

func DepartamentosConsultados(c *gin.Context) {
	departamentos := []Departamento{
		Departamento{
			Direccion:    "Sitio al 1200",
			Inmobiliaria: "San Miguel",
		},
		Departamento{
			Direccion:    "Las Piedras al 1500",
			Inmobiliaria: "German algo",
		},
	}
	c.JSON(http.StatusOK, departamentos)
}

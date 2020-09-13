package routes

import (
	"braiton/braiton-home/cmd/departmentsearch/routes/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes() {
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

	router.GET("api/departamentos-consultados", handlers.DepartmentConsulted)

	router.Run(":" + port)
}

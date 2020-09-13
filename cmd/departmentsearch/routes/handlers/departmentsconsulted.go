package handlers

import (
	"braiton/braiton-home/cmd/internal/departments"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DepartmentConsulted(c *gin.Context) {
	departamentos := departments.GetDepartments()
	c.JSON(http.StatusOK, departamentos)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/models"
	"github.com/koreset/empkore/services"
	"net/http"
	"strconv"
)

func CreateEmployee(c *gin.Context) {
	var emp models.Employee
	err := c.ShouldBind(&emp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An unspecified error occurred with the request"})
		return
	}

	if isValidEmployee(emp) {
		services.CreateNewEmployee(&emp)
		c.Header("Content-Location", "/employees/"+strconv.Itoa(int(emp.ID)))
		c.JSON(http.StatusCreated, emp)

	}else{
		c.JSON(http.StatusBadRequest, gin.H{"error":"There were one or more validation errors encountered"})
	}

}

func isValidEmployee(emp models.Employee) bool {
	if len(emp.FirstName) < 1{
		return  false
	}

	if len(emp.LastName) < 1{
		return false
	}

	if len(emp.Email) < 1 {
		return false
	}

	if len(emp.Password) < 1 {
		return false
	}

	return true
}

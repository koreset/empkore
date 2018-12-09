package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/models"
)

var err error

func Login(c *gin.Context) {

}

func registerNewEmployee(emp models.Employee) (*models.Employee, error) {
	return nil, err
}

func authenticateEmployee(email, password string) bool {
	return false
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var err error

type UserLogin struct {
	Email    string `json:"email" form:"password"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var userLogin UserLogin

	err = c.ShouldBind(&userLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something is not quite right with the request"})
		return
	}

	if isValid(userLogin) {
		if authenticateEmployee(userLogin.Email, userLogin.Password) {
			c.JSON(http.StatusOK, gin.H{"token": "123456789"})
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Supplied credentials are invalid"})
}


func authenticateEmployee(email, password string) bool {
	foundEmployee, err := services.GetEmployeeByEmail(email)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundEmployee.Password), []byte(password))

	if err == nil {
		return true
	}

	return false
}

func isValid(userLogin UserLogin) bool {
	if len(userLogin.Email) > 0 && len(userLogin.Password) > 0 {
		return true
	}
	//We should run more validations here later

	return false
}

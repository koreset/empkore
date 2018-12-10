package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var err error

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context) {

	if c.Request.Method == http.MethodGet {
		title := "Login"
		c.HTML(http.StatusOK, "employee-login", gin.H{"title": title})
	} else {
		var userLogin UserLogin
		err = c.ShouldBind(&userLogin)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Something is not quite right with the request"})
			return
		}

		if isValid(userLogin) {
			if authenticateEmployee(userLogin.Email, userLogin.Password) {
				if c.Request.Header.Get("Accept") == "application/json" {
					c.JSON(http.StatusOK, gin.H{"token": "123456789"})
					return
				} else {

					c.Redirect(http.StatusMovedPermanently, "/")
					//c.HTML(http.StatusOK, "employee-index", gin.H{"is_logged_on": true})
					return
				}
			}
		}
		//render(c, http.StatusUnauthorized, gin.H{"error": "Supplied credentials are invalid"}, "" )
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Supplied credentials are invalid"})
	}
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

func render(c *gin.Context, status int, data interface{}, template string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(status, data)
	default:
		c.HTML(status, template, data)
	}
}

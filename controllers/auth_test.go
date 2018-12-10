package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/services"
	"github.com/koreset/empkore/utils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gin-gonic/gin.v1/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthenticateEmployee(t *testing.T) {
	//We need to do authentication for both api and front end access.
	//So a common authentication function should be written, right?

	//Given
	setupDB()
	testEmployee := utils.GetValidEmployee()
	testPassword := testEmployee.Password
	services.CreateNewEmployee(&testEmployee)

	//When
	result := authenticateEmployee(testEmployee.Email, testPassword)

	//Then
	assert.Equal(t, true, result)
	cleanupDB()

}

func TestLogin(t *testing.T) {
	//Given
	setupDB()
	var userLogin = UserLogin{
		Email:    "jome@koreset.com",
		Password: "wordpass15",
	}

	testEmp := utils.GetValidEmployee()

	services.CreateNewEmployee(&testEmp)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/login", Login)
	payload, _ := json.Marshal(userLogin)
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(string(payload)))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	//When
	r.ServeHTTP(w, req)

	//Then
	assert.Equal(t, http.StatusOK, w.Code)
	body, _ := ioutil.ReadAll(w.Body)
	assert.Contains(t, string(body), `{"token":"123456789"}`)

	//Finally
	cleanupDB()

}

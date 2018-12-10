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
	"net/url"
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

func TestLoginGetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	templates, _ := utils.SetupTemplates("../views")
	r.SetHTMLTemplate(templates)
	r.GET("/login", Login)
	req, _ := http.NewRequest("GET", "/login", nil)

	r.ServeHTTP(w, req)

	content, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, string(content), "<title>Login</title>")
	assert.Contains(t, string(content), `<label for="email">Email address</label>`)
	assert.Contains(t, string(content), `<label for="password">Password</label>`)

}

func TestLoginViaAPI(t *testing.T) {
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


func TestLoginViaForm(t *testing.T){
	setupDB()
	testEmp := utils.GetValidEmployee()
	services.CreateNewEmployee(&testEmp)

	r := gin.Default()
	r.POST("/login", Login)
	templates, _ := utils.SetupTemplates("../views")
	r.SetHTMLTemplate(templates)

	params := url.Values{}
	params.Add("email", "jome@koreset.com")
	params.Add("password","wordpass15")

	payload := params.Encode()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	content, _ := ioutil.ReadAll(w.Body)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, string(content), "Dashboard")
	cleanupDB()
}
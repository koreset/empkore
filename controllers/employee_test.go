package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/models"
	"github.com/koreset/empkore/services"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	code := m.Run()

	setupDB()

	services.GetDB().DropTableIfExists(&models.Employee{}, &models.Position{})

	os.Exit(code)
}

//TestCreateEmployee will test the creation of a new employee object
func TestCreateEmployeeFormPayload(t *testing.T) {
	//Given
	employeePayload := getEmployeeFormPayload()

	var actual models.Employee
	r := getRouter(false)
	r.POST("/employees/new", CreateEmployee)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/employees/new", strings.NewReader(employeePayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(employeePayload)))

	//When
	r.ServeHTTP(w, req)

	//Then
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "/employees/1", w.Header().Get("Content-Location"))
	content, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(content, &actual)

	assert.Equal(t, 1, int(actual.ID))
	assert.Equal(t, "Jome", actual.FirstName)
	cleanupDB()

}

func TestCreateEmployeeJsonPayload(t *testing.T) {
	//Given
	setupDB()
	var actual models.Employee
	newEmp := getValidEmployee()
	newEmpJson, _ := json.Marshal(newEmp)

	r := getRouter(false)
	r.POST("/employees/new", CreateEmployee)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST","/employees/new", strings.NewReader(string(newEmpJson)))
	req.Header.Add("Content-Type", "application/json")

	//When
	r.ServeHTTP(w, req)

	//Then
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "/employees/1", w.Header().Get("Content-Location"))
	content, _ := ioutil.ReadAll(w.Body)

	json.Unmarshal(content, &actual)

	assert.Equal(t, 1, int(actual.ID))
	assert.Equal(t, "Jome", actual.FirstName)
	cleanupDB()
}

func TestCreateEmployeeThrowInvalidErrors(t *testing.T) {
	//Given
	incompletePayload := getIncompleteEmployeeFormPayload()
	r := getRouter(false)
	r.POST("/employees/new", CreateEmployee)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/employees/new", strings.NewReader(incompletePayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(incompletePayload)))
	//When
	r.ServeHTTP(w, req)
	//Then
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestIsValidEmployee(t *testing.T) {
	newEmp := getValidEmployee()
	assert.True(t, isValidEmployee(newEmp))

	invalidEmp := models.Employee{}

	assert.False(t, isValidEmployee(invalidEmp))
}

func getValidEmployee() models.Employee {
	newEmp := models.Employee{
		ID:        1,
		FirstName: "Jome",
		LastName:  "Akpoduado",
		Email:     "jome@koreset.com",
		CellPhone: "0719166815",
		JoinDate:  time.Date(2018, time.January, 01, 0, 0, 0, 0, time.UTC),
		Password:  "wordpass15",
	}
	return newEmp
}

func getEmployeeFormPayload() string {
	formParams := url.Values{}
	formParams.Add("first_name", "Jome")
	formParams.Add("last_name", "Akpoduado")
	formParams.Add("email", "jome@koreset.com")
	formParams.Add("password", "wordpass15")
	formParams.Add("join_date", "2018-01-01")
	formParams.Add("cell_phone", "0719166815")

	return formParams.Encode()
}

func getIncompleteEmployeeFormPayload() string {
	formParams := url.Values{}
	formParams.Add("first_name", "")
	formParams.Add("last_name", "")
	formParams.Add("email", "")
	formParams.Add("password", "")
	formParams.Add("join_date", "")
	formParams.Add("cell_phone", "")

	return formParams.Encode()
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../views")
	}

	return r
}

func cleanupDB(){
	services.GetDB().DropTableIfExists(&models.Employee{}, &models.Position{})
}

func setupDB(){
	services.GetDB().AutoMigrate(&models.Position{}, &models.Employee{})
}
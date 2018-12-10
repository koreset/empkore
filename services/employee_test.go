package services

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/models"
	"github.com/koreset/empkore/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testing"
	"time"
)

var tmpEmployeeList []models.Employee

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	code := m.Run()

	os.Exit(code)
}

func TestGetAllEmployees(t *testing.T) {
	tmpEmployeeList := GetAllEmployees()

	assert.Equal(t, 2, len(tmpEmployeeList))

	tmpEmployeeList = []models.Employee{}

}

func TestCreateNewEmployee(t *testing.T) {
	setupDB()
	var newEmployee = utils.GetValidEmployee()
	var expectedEmployee = utils.GetValidEmployee()

	CreateNewEmployee(&newEmployee)

	assert.Equal(t, expectedEmployee.FirstName, newEmployee.FirstName)
	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(newEmployee.Password), []byte(expectedEmployee.Password)))

	cleanupDb()

}

func TestGetEmployeeByID(t *testing.T) {
	tmpEmployeeList = GetAllEmployees()
	expectedEmployee := models.Employee{ID: 1, FirstName: "Jome", MiddleName: "Christopher", LastName: "Akpoduado", Email: "jome@koreset.com", JoinDate: time.Date(2017, time.March, 01, 0, 0, 0, 0, time.UTC)}
	actualEmployee, _ := GetEmployeeByID(1)

	assert.Equal(t, expectedEmployee, actualEmployee)
}

func TestEncryptPassword(t *testing.T) {
	// Given
	password := "+w3ak15-@BlaqBee04"

	//When
	hashed := encryptPassword(password)

	//Then
	assert.Equal(t, 60, len(hashed))
	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)))
}

func setupDB() {
	GetDB().DropTableIfExists(&models.Employee{}, models.Position{})
	GetDB().AutoMigrate(&models.Employee{}, models.Position{})
}
func cleanupDb() {
	GetDB().DropTable(&models.Employee{})
}

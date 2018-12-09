package services

import (
	"github.com/koreset/empkore/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var tmpEmployeeList []models.Employee

func TestGetAllEmployees(t *testing.T) {
	tmpEmployeeList := GetAllEmployees()

	assert.Equal(t, 2, len(tmpEmployeeList))

	tmpEmployeeList = []models.Employee{}

}

func TestCreateNewEmployee(t *testing.T) {

	var newEmployee = models.Employee{ID: 3, FirstName: "Kuzo", LastName: "Dasa", Email: "kuzo@koreset.com", JoinDate: time.Date(2018, time.April, 01, 0, 0, 0, 0, time.UTC)}
	CreateNewEmployee(&newEmployee)
	tmpEmployeeList = GetAllEmployees()

	assert.Equal(t, 3, len(tmpEmployeeList))
	assert.Contains(t, tmpEmployeeList, newEmployee)
	tmpEmployeeList = []models.Employee{}

	cleanupDb()

}

func TestGetEmployeeByID(t *testing.T) {
	tmpEmployeeList = GetAllEmployees()
	expectedEmployee := models.Employee{ID: 1, FirstName: "Jome", MiddleName: "Christopher", LastName: "Akpoduado", Email: "jome@koreset.com", JoinDate: time.Date(2017, time.March, 01, 0, 0, 0, 0, time.UTC)}
	actualEmployee, _ := GetEmployeeByID(1)

	assert.Equal(t, expectedEmployee, actualEmployee)
}

func cleanupDb() {
	GetDB().DropTable(&models.Employee{})
}

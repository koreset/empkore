package services

import (
	"github.com/koreset/empkore/models"
	"time"
)

var employeeList = []models.Employee{
	{ID: 1, FirstName: "Jome", MiddleName: "Christopher", LastName: "Akpoduado", Email: "jome@koreset.com", JoinDate: time.Date(2017, time.March, 01, 0, 0, 0, 0, time.UTC)},
	{ID: 2, FirstName: "Emile", MiddleName: "Charles", LastName: "Senga", Email: "emile@koreset.com", JoinDate: time.Date(2017, time.August, 01, 0, 0, 0, 0, time.UTC)},
}

func CreateNewEmployee(employee *models.Employee) {
	GetDB().Create(&employee)
	employeeList = append(employeeList, *employee)
}

func GetAllEmployees() []models.Employee {
	return employeeList
}

func GetEmployeeByID(id uint) (models.Employee, error) {
	for _, emp := range employeeList {
		if emp.ID == id {
			return emp, nil
		}
	}
	return models.Employee{}, nil
}

package services

import (
	"github.com/koreset/empkore/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var employeeList = []models.Employee{
	{ID: 1, FirstName: "Jome", MiddleName: "Christopher", LastName: "Akpoduado", Email: "jome@koreset.com", JoinDate: time.Date(2017, time.March, 01, 0, 0, 0, 0, time.UTC)},
	{ID: 2, FirstName: "Emile", MiddleName: "Charles", LastName: "Senga", Email: "emile@koreset.com", JoinDate: time.Date(2017, time.August, 01, 0, 0, 0, 0, time.UTC)},
}

func CreateNewEmployee(employee *models.Employee) {
	employee.Password = encryptPassword(employee.Password)
	GetDB().Create(&employee)
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

func GetEmployeeByEmail(email string) (models.Employee, error) {
	var emp models.Employee

	err := GetDB().Where("email = ?", email).First(&emp).Error
	if err != nil{
		return emp, err
	}
	return emp, nil
}

func encryptPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

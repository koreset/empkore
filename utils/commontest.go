package utils

import (
	"github.com/koreset/empkore/models"
	"time"
)

func GetValidEmployee() models.Employee {
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

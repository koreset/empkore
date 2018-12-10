package models

import (
	"time"
)

type Employee struct {
	ID         uint      `form:"id" json:"id" gorm:"primary_key;auto_increment"`
	FirstName  string    `form:"first_name" json:"first_name"`
	MiddleName string    `form:"middle_name" json:"middle_name"`
	LastName   string    `form:"last_name" json:"last_name"`
	Email      string    `form:"email" json:"email" gorm:"unique"`
	CellPhone  string    `form:"cell_phone" json:"cell_phone"`
	JoinDate   time.Time `form:"join_date" time_format:"2006-01-02" time_utc:"2" json:"join_date"`
	Password   string    `form:"password" json:"password"`
	PositionID uint
	Position   Position
}

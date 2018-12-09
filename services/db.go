package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/koreset/empkore/models"
)

var db *gorm.DB
var err error

func InitializeDB() error {
	if gin.Mode() == gin.TestMode {
		db, err = gorm.Open("sqlite3", "../empkoretest.db")
	} else {
		db, err = gorm.Open("mysql", "root:wordpass15@tcp(localhost:3306)/empkoredb?charset=utf8&parseTime=True&loc=local")
	}

	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Employee{}, &models.Position{})

	var employee models.Employee
	var position models.Position

	db.Model(&position).Related(&employee)

	return nil
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	if err = InitializeDB(); err != nil {
		return nil
	}

	return db
}

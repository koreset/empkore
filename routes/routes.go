package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/controllers"
)

func InitializeRoutes(r *gin.Engine){
	r.GET("/", controllers.Home)
	r.POST("/login", controllers.Login)
	r.POST("/employees/new", controllers.CreateEmployee)

	r.Static("/public", "./assets")
}

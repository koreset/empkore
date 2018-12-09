package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/controllers"
)

func InitializeRoutes(router *gin.Engine){
	router.GET("/", controllers.Home)
	router.Static("/public", "./assets")
}

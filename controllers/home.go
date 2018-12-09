package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context){
	c.HTML(http.StatusOK, "employee-index", nil)
}

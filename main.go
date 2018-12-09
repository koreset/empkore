package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/routes"
	"github.com/koreset/empkore/utils"
	"log"
)

var router *gin.Engine

func main() {

	router := gin.Default()
	templates, err := utils.SetupTemplates("views")

	if err != nil{
		fmt.Println("We should be aborting now..")
	}

	router.SetHTMLTemplate(templates)

	routes.InitializeRoutes(router)

	log.Fatal(router.Run())
}

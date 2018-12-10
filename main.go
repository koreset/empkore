package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koreset/empkore/routes"
	"github.com/koreset/empkore/services"
	"github.com/koreset/empkore/utils"
	"log"
)

var router *gin.Engine

func main() {
	err:= services.InitializeDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	router := gin.Default()
	templates, err := utils.SetupTemplates("views")

	if err != nil {
		fmt.Println("We should be aborting now..")
	}

	router.SetHTMLTemplate(templates)

	routes.InitializeRoutes(router)

	log.Fatal(router.Run())
}

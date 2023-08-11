package main

import (
	"RESTAPI/constants"
	"RESTAPI/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//router.GET("/") this get is in routes file
	routes.AppRoutes(router) // we write this
	routes.AppRoutesP(router)
	fmt.Println("server runnign on port", constants.Port)
	log.Fatal(router.Run(constants.Port))
}

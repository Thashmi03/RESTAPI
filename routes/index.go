package routes

import (
	"RESTAPI/controllers"

	"github.com/gin-gonic/gin"
)
//paramter is pointer ref to gin engine
func AppRoutes(router *gin.Engine) {
	router.GET("/api/login", controllers.Login)
	// router.POST("/api/register")
	// router.GET("/api/profile:id")
}
package routes

import (
	"RESTAPI/controllers"

	"github.com/gin-gonic/gin"
)
//paramter is pointer ref to gin engine
func AppRoutes(router *gin.Engine) {
	router.POST("/api/login/name", controllers.Login)
	router.GET("/api/register",controllers.Register)
	// router.GET("/api/profile:id")
}


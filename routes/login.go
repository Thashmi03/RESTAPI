package routes

import (
	"RESTAPI/controllers"

	"github.com/gin-gonic/gin"
)
//paramter is pointer ref to gin engine
func AppRoutesP(router *gin.Engine) {
	router.GET("/api/loginp", controllers.LoginP)
	// router.POST("/api/register")
	// router.GET("/api/profile:id")
}
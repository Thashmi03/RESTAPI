package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Login(c * gin.Context){
	// name:=c.Param("name")
	c.String(http.StatusOK,"hello")
}
// func Register(c * gin.Context){

// }
// func GetProfile(c * gin.Context){

// }
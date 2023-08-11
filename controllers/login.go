package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func LoginP(c * gin.Context){
	// name:=c.Param("name")
	c.String(http.StatusOK,"hellologin")
}
// func Register(c * gin.Context){

// }
// func GetProfile(c * gin.Context){

// }
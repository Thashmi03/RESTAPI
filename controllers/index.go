package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s", name)
}
func Register(c *gin.Context) {
	c.String(http.StatusOK, "registered")
}

// func GetProfile(c * gin.Context){

// }

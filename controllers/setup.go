package controllers

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetRouting() {
	router = gin.Default()
}

func SetLoginRoutes() {
	loginRoutes := router.Group("/api/login")
	{
		loginRoutes.POST("/", Login)
	}
}

func Run() {
	router.Run()
}

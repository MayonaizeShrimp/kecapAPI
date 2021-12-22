package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetRouting() {
	router = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
}

func SetLoginRoutes() {
	loginRoutes := router.Group("/api/login")
	{
		loginRoutes.POST("/", Login)
	}
}

func SetAPIRoutes() {
	api := router.Group("/api/kecap")
	{
		api.GET("/", getAll)
		api.GET("/:id", get)
		api.POST("/", insert)
		api.PATCH("/", update)
		api.DELETE("/", delete)
	}
}

func Run() {
	router.Run()
}

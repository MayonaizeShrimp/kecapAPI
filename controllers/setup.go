package controllers

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//Initiate Gin
func SetRouting() {
	router = gin.Default()
}

//Set routes for login
func SetLoginRoutes() {
	loginRoutes := router.Group("/api/login")
	{
		loginRoutes.POST("/", Login)
	}
}

//set routes for API
func SetAPIRoutes() {
	api := router.Group("/api/kecap")
	{
		api.GET("/", getAll)    //select all
		api.GET("/:id", get)    //select by id
		api.POST("/", insert)   //create
		api.PATCH("/", update)  //update
		api.DELETE("/", delete) //delete
	}
}

//Execute Gin in port 8080
func Run() {
	router.Run()
}

package controllers

import (
	"time"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//Initiate Gin
func SetRouting() {
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Run()
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

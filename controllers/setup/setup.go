package controllerssetup

import (
	"kecapstore/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetRouting() {
	router = gin.Default()
}

func SetLoginRoutes() {
	loginRoutes := router.Group("/api/login")
	{
		loginRoutes.GET("/", controllers.Login)
	}
}

func Run() {
	router.Run()
}

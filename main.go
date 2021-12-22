package main

import (
	"kecapstore/controllers"
	"kecapstore/models"
)

func main() {
	models.ConnectDatabase()

	controllers.SetRouting()
	controllers.SetLoginRoutes()
	controllers.SetAPIRoutes()
	controllers.Run()
}

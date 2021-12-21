package main

import (
	controllerssetup "kecapstore/controllers/setup"
	modelssetup "kecapstore/models/setup"
)

func main() {
	modelssetup.ConnectDatabase()

	controllerssetup.SetRouting()
	controllerssetup.SetLoginRoutes()
	controllerssetup.Run()
}

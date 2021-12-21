package controllerssetup

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

func Run() {
	router.Run()
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kecapstore/models"
)

func Login(c *gin.Context) {
	var loginData models.LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginData.Verify() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user found with that email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": loginData})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kecapstore/models"
)

func Login(c *gin.Context) {
	var loginData models.LoginData

	//check if user input data types fields
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//verify login data with database
	if loginData.Verify() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No user found with that email or password"})
		return
	}

	//return login data
	c.JSON(http.StatusOK, gin.H{"data": loginData})
}

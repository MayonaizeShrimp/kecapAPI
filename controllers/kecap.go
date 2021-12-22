package controllers

import (
	"kecapstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAll(c *gin.Context) {
	var kecaps []models.Kecap
	models.DB.Find(&kecaps)

	c.JSON(http.StatusOK, gin.H{"data": kecaps})
}

func get(c *gin.Context) {
	var kecap models.Kecap

	//find kecap by id
	if err := models.DB.Where("id = ?", c.Param("id")).First(&kecap).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kecap})
}

func insert(c *gin.Context) {
	var input models.KecapInsertData

	//check if user input is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//insert to db
	kecap := models.Kecap{Nama: input.Nama, Harga: input.Harga}
	models.DB.Create(&kecap)

	c.JSON(http.StatusOK, gin.H{"data": kecap})
}

func update(c *gin.Context) {
	var input models.KecapUpdateData

	//check if user input is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if record exist
	var kecap models.Kecap
	if err := models.DB.Where("id = ?", c.Param("id")).First(&kecap).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//try udpate the row
	kecap = models.Kecap{Nama: input.Nama, Harga: input.Harga}
	models.DB.Model(&kecap).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": kecap})
}

func delete(c *gin.Context) {
	var kecap models.Kecap

	//check if record exist
	if err := models.DB.Where("id = ?", c.Param("id")).First(&kecap).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//try to delete
	models.DB.Delete(&kecap)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

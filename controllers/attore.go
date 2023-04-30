package controllers

import (
	"cdbp4/initializers"
	"cdbp4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAttori(c *gin.Context) {
	var attori []models.Attore
	initializers.DB.Find(&attori)

	data := gin.H{
		"attori": attori,
	}

	c.HTML(http.StatusOK, "Attori.html", data)

}

func CreateAttore(c *gin.Context) {
	var attore models.Attore

	attore.Nome = c.PostForm("nome")
	attore.Cognome = c.PostForm("cognome")
	attore.Email = c.PostForm("email")

	if err := initializers.DB.Create(&attore).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadAttori(c)
}

func DeleteAttore(c *gin.Context) {
	id := c.Param("id")

	var attore models.Attore
	if err := initializers.DB.First(&attore, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Teatro non trovato"})
		return
	}

	if err := initializers.DB.Delete(&attore).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadAttori(c)
}

func GetUpdateAttore(c *gin.Context) {
	var attore models.Attore
	id := c.Param("id")
	if err := initializers.DB.First(&attore, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attore non trovato"})
	}

	data := gin.H{
		"attore": attore,
	}

	c.HTML(http.StatusOK, "Attore.html", data)
}

func PostUpdateAttore(c *gin.Context) {
	id := c.Param("id")

	var attore models.Attore
	if err := initializers.DB.First(&attore, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Teatro non trovato"})
		return
	}

	attore.Nome = c.PostForm("nome")
	attore.Cognome = c.PostForm("cognome")
	attore.Email = c.PostForm("email")

	if err := initializers.DB.Save(&attore).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadAttori(c)
}

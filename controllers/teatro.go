package controllers

import (
	"cdbp4/initializers"
	"cdbp4/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadTeatri(c *gin.Context) {
	var teatri []models.Teatro
	initializers.DB.Find(&teatri)

	data := gin.H{
		"teatri": teatri,
	}

	c.HTML(http.StatusOK, "Teatri.html", data)

}

func CreateTeatri(c *gin.Context) {
	var teatro models.Teatro
	if err := c.ShouldBind(&teatro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teatro.Nome = c.PostForm("nome")
	postiForm := c.PostForm("posti")
	posti, err := strconv.Atoi(postiForm)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	teatro.Posti = int(posti)

	if err := initializers.DB.Create(&teatro).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadTeatri(c)
}

func DeleteTeatri(c *gin.Context) {
	id := c.Param("id")

	var teatro models.Teatro
	if err := initializers.DB.First(&teatro, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Teatro non trovato"})
		return
	}

	if err := initializers.DB.Delete(&teatro).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadTeatri(c)
}

func GetUpdateTeatri(c *gin.Context) {
	var teatro models.Teatro
	id := c.Param("id")
	if err := initializers.DB.First(&teatro, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teatro non trovato"})
	}

	data := gin.H{
		"teatro": teatro,
	}

	c.HTML(http.StatusOK, "Teatro.html", data)
}

func PostUpdateTeatri(c *gin.Context) {
	id := c.Param("id")

	var teatro models.Teatro
	if err := initializers.DB.First(&teatro, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Teatro non trovato"})
		return
	}

	if err := c.ShouldBind(&teatro); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	teatro.Nome = c.PostForm("nome")
	postiForm := c.PostForm("posti")
	posti, err := strconv.Atoi(postiForm)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	teatro.Posti = int(posti)

	if err := initializers.DB.Save(&teatro).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadTeatri(c)
}

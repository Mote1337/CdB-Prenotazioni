package controllers

import (
	"cdbp4/initializers"
	"cdbp4/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// READ BIGLIETTO
func ReadBiglietto(c *gin.Context) {
	// Creazione slice che ospiterà tutte le entry da db
	var biglietti []models.Biglietto
	// Assegnazioni entry da tabella allo struct
	initializers.DB.Order(c.DefaultQuery("orderby", "ID")).Find(&biglietti)

	// Creo i dati da passare alla pagina html
	data := gin.H{
		"biglietti": biglietti,
	}
	// Render della pagina con i dati passati
	c.HTML(http.StatusOK, "Biglietti.html", data)
}

// CREATE BIGLIETTO
func CreateBiglietto(c *gin.Context) {
	var biglietti models.Biglietto
	if err := c.ShouldBind(&biglietti); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	biglietti.Tipo = c.PostForm("tipo")
	costoForm := c.PostForm("costo")
	costo, err := strconv.Atoi(costoForm)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	biglietti.Costo = int(costo)

	if err := initializers.DB.Create(&biglietti).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadBiglietto(c)
}

func DeleteBiglietto(c *gin.Context) {
	id := c.Param("id")

	var biglietto models.Biglietto
	if err := initializers.DB.First(&biglietto, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Biglietto non trovato"})
		return
	}

	if err := initializers.DB.Delete(&biglietto).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadBiglietto(c)
}

func GetUpdateBiglietto(c *gin.Context) {
	var biglietto models.Biglietto
	id := c.Param("id")
	if err := initializers.DB.First(&biglietto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Biglietto non trovato"})
	}

	data := gin.H{
		"biglietto": biglietto,
	}

	c.HTML(http.StatusOK, "Biglietto.html", data)
}

func PostUpdatBiglietto(c *gin.Context) {
	id := c.Param("id")

	var biglietto models.Biglietto
	if err := initializers.DB.First(&biglietto, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Biglietto non trovato"})
		return
	}

	if err := c.ShouldBind(&biglietto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	biglietto.Tipo = c.PostForm("tipo")
	costoForm := c.PostForm("costo")
	costo, err := strconv.Atoi(costoForm)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	biglietto.Costo = int(costo)

	if err := initializers.DB.Save(&biglietto).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadBiglietto(c)
}

package controllers

import (
	"cdbp4/helper"
	"cdbp4/initializers"
	"cdbp4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// READ SPETTACOLO
func ReadSpettacoli(c *gin.Context) {
	// Creazione slice che ospiterà tutte le entry da db
	var spettacoli []models.Spettacolo
	// Assegnazioni entry da tabella allo struct
	initializers.DB.Order(c.DefaultQuery("orderby", "ID")).Find(&spettacoli)

	var teatri []models.Teatro
	initializers.DB.Find(&teatri)
	// Creo i dati da passare alla pagina html
	data := gin.H{
		"spettacoli": spettacoli,
		"teatri":     teatri,
	}
	// Render della pagina con i dati passati
	c.HTML(http.StatusOK, "Spettacoli.html", data)
}

// CREATE SPETTACOLO
func CreateSpettacolo(c *gin.Context) {
	var spettacolo models.Spettacolo
	if err := c.ShouldBind(&spettacolo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	spettacolo.Nome = c.PostForm("nome")

	spettacolo.Inizio = helper.ConvertiData(c.PostForm("inizio"))

	spettacolo.Fine = helper.ConvertiData(c.PostForm("fine"))

	spettacolo.TeatroID = helper.ReturnUIntFromString(c, c.PostForm("teatro"))

	var teatro models.Teatro
	initializers.DB.First(&teatro, c.PostForm("teatro"))
	spettacolo.TeatroName = teatro.Nome

	if err := initializers.DB.Create(&spettacolo).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadSpettacoli(c)
}

// UPDATE SPETTACOLO
// GET
func GetUpdateSpettacolo(c *gin.Context) {
	id := c.Param("id")

	var spettacolo models.Spettacolo
	if err := initializers.DB.First(&spettacolo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Spettacolo non trovato"})
		return
	}

	var teatri []models.Teatro
	if err := initializers.DB.Find(&teatri).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teatro Selezionato non trovato"})
	}

	data := gin.H{
		"spettacolo": spettacolo,
		"teatri":     teatri,
	}
	c.HTML(http.StatusOK, "Spettacolo.html", data)
}

// POST
func PostUpdateSpettacolo(c *gin.Context) {
	id := c.Param("id")

	var spettacolo models.Spettacolo
	if err := initializers.DB.First(&spettacolo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Spettacolo non trovato"})
		return
	}

	spettacolo.Nome = c.PostForm("nome")

	inizioStr := c.PostForm("inizio")
	spettacolo.Inizio = helper.ConvertiData(inizioStr)

	fineStr := c.PostForm("fine")
	spettacolo.Fine = helper.ConvertiData(fineStr)

	teatroid := c.PostForm("teatro")
	spettacolo.TeatroID = helper.ReturnUIntFromString(c, teatroid)

	var teatro models.Teatro
	initializers.DB.First(&teatro, teatroid)
	spettacolo.TeatroName = teatro.Nome

	if err := initializers.DB.Save(&spettacolo).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadSpettacoli(c)
}

// DELETE SPETTACOLO
func DeleteSpettacolo(c *gin.Context) {
	id := c.Param("id")

	var spettacolo models.Spettacolo
	if err := initializers.DB.First(&spettacolo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Spettacolo non trovato"})
		return
	}

	if err := initializers.DB.Delete(&spettacolo).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ReadSpettacoli(c)
}

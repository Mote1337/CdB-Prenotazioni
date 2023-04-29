package controllers

import (
	"cdbp4/helper"
	"cdbp4/initializers"
	"cdbp4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadPrenotazioniUnico(c *gin.Context) {
	// Creazione struct che ospiterà la determinata entry da DB
	var spettacolo models.Spettacolo
	// Reperisco l'id per la query
	id := c.Param("id")
	// Assegnazione della entry da DB allo struct con condizione della chiave primaria id
	if err := initializers.DB.First(&spettacolo, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Spettacolo non trovato"})
	}
	// Creazione slice di date
	date := helper.DateRange(spettacolo.Inizio, spettacolo.Fine)
	// Numero minimo e massimo di prenotazioni per singola persona
	numeriPosti := helper.SlicePosti(1, 10)

	dateQuery := c.DefaultQuery("data", "NODATA")

	if dateQuery == "NODATA" {
		var prenotazioni []models.Prenotazione
		initializers.DB.Where("spettacolo_id = ?", spettacolo.ID).Order(c.Query("orderby")).Find(&prenotazioni)

		data := gin.H{
			"prenotazioni": prenotazioni,
			"spettacolo":   spettacolo,
			"date":         date,
			"numeriPosti":  numeriPosti,
		}

		c.HTML(http.StatusOK, "Prenotazioni.html", data)
	} else {
		var prenotazioni []models.Prenotazione
		initializers.DB.Where("data::date = ?", dateQuery).Order(c.Query("orderby")).Find(&prenotazioni)

		data := gin.H{
			"prenotazioni":   prenotazioni,
			"spettacolo":     spettacolo,
			"date":           date,
			"dateQuery":      dateQuery,
			"numeriPosti":    numeriPosti,
			"postiprenotati": helper.GetPostiPrenotati(dateQuery),
		}
		c.HTML(http.StatusOK, "PrenotazioniGiorno.html", data)
	}

}

func PostCreatePrenotazione(c *gin.Context) {
	var prenotazione models.Prenotazione

	// Campo Prenotazione Nome
	prenotazione.Nome = c.PostForm("nome")

	// Campo Prenotazione Posti
	prenotazione.Posti = helper.ReturnIntFromString(c, c.PostForm("posti"))

	// Campo Prenotazione Data
	prenotazione.Data = helper.ConvertiData(c.PostForm("data"))

	// Campo Prenotazione SpettacoloID
	prenotazione.SpettacoloID = helper.ReturnUIntFromString(c, c.PostForm("spettacoloID"))

	// Campo Prenotazione SpettacoloNome
	var spettacolo models.Spettacolo
	initializers.DB.First(&spettacolo, prenotazione.SpettacoloID)
	prenotazione.SpettacoloName = spettacolo.Nome

	// Campo Prenotazione Referente
	prenotazione.Referente = c.PostForm("referente")

	// Creazione Prenotazione con i dati reperiti
	if err := initializers.DB.Create(&prenotazione).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Logica per tornare alla prenotazione

	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: c.PostForm("spettacoloID"),
		},
	}

	if c.Query("data") != "" {
		c.Request.URL.Query().Add("data", c.Query("data"))
	}

	ReadPrenotazioniUnico(c)
}

func GetUpdatePrenotazione(c *gin.Context) {
	// Ricavo ID prenotazione per effettuare la ricerca sul db esatta per record
	id := c.Param("id")
	// Mi creo la variabile struct dove inserire il record da DB
	var prenotazione models.Prenotazione
	// Interrogo il DB assegnando la SELECT risultante alla variabile prenotazione
	if err := initializers.DB.First(&prenotazione, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Prenotazione non trovata"})
		return
	}

	// Mi genero l'array per i posti minimo e massimi prenotabili
	slicePosti := helper.SlicePosti(1, 10)

	// Mi creo lo struct Spettacolo per ricavare le date
	var spettacolo models.Spettacolo
	if err := initializers.DB.First(&spettacolo, prenotazione.SpettacoloID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Spettacolo non trovato"})
		return
	}

	// Mi ricavo le date disponibili
	date := helper.DateRange(spettacolo.Inizio, spettacolo.Fine)

	data := gin.H{
		"prenotazione": prenotazione,
		"numeriPosti":  slicePosti,
		"date":         date,
		"spettacolo":   spettacolo,
	}
	c.HTML(http.StatusOK, "Prenotazione.html", data)
}

func PostUpdatePrenotazione(c *gin.Context) {
	id := c.Param("id")

	var prenotazione models.Prenotazione
	if err := initializers.DB.First(&prenotazione, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Spettacolo non trovato"})
		return
	}

	// Campo Prenotazione Nome
	prenotazione.Nome = c.PostForm("nome")

	// Campo Prenotazione Posti
	prenotazione.Posti = helper.ReturnIntFromString(c, c.PostForm("posti"))

	// Campo Prenotazione Data
	prenotazione.Data = helper.ConvertiData(c.PostForm("data"))

	// Campo Prenotazione SpettacoloID
	prenotazione.SpettacoloID = helper.ReturnUIntFromString(c, c.PostForm("spettacoloID"))

	// Campo Prenotazione SpettacoloName Rimane uguale

	// Campo Prenotazione Referente
	prenotazione.Referente = c.PostForm("referente")

	// Aggiornamento Prenotazione con i dati reperiti
	if err := initializers.DB.Save(&prenotazione).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Logica per tornare alla prenotazione dello spettacolo
	idSpettacolo := strconv.FormatUint(uint64(prenotazione.SpettacoloID), 10)
	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: idSpettacolo,
		},
	}

	ReadPrenotazioniUnico(c)
}

func DeletePrenotazione(c *gin.Context) {
	id := c.Param("id")

	var prenotazione models.Prenotazione
	if err := initializers.DB.First(&prenotazione, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Prenotazione non trovata"})
		return
	}

	idSpettacolo := strconv.FormatUint(uint64(prenotazione.SpettacoloID), 10)

	if err := initializers.DB.Delete(&prenotazione).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: idSpettacolo,
		},
	}

	ReadPrenotazioniUnico(c)
}

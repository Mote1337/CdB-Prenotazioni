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
	// Get Lista attori
	var listaAttori []models.Attore
	if err := initializers.DB.Find(&listaAttori).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attori non trovati"})
	}
	// Get Biglietti
	var biglietti []models.Biglietto
	if err := initializers.DB.Find(&biglietti).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Biglietti non trovati"})
	}
	// Calcolo totale dei biglietti
	bigliettiPerTipo := make(map[string]struct {
		NumBiglietti int
		CostoTotale  int
	})

	dateQuery := c.DefaultQuery("data", "NODATA")

	if dateQuery == "NODATA" {
		var prenotazioni []models.Prenotazione
		initializers.DB.Where("spettacolo_id = ?", spettacolo.ID).Order(c.Query("orderby")).Find(&prenotazioni)
		bigliettiPerTipo = helper.BigliettiTotali(prenotazioni)

		var totaleBiglietti, totaleCosti int
		for _, val := range bigliettiPerTipo {
			totaleBiglietti += val.NumBiglietti
			totaleCosti += val.CostoTotale
		}

		data := gin.H{
			"prenotazioni":         prenotazioni,
			"spettacolo":           spettacolo,
			"date":                 date,
			"numeriPosti":          numeriPosti,
			"listaAttori":          listaAttori,
			"postiTotaliPrenotati": helper.PostiTotaliPrenotati(prenotazioni),
			"postiTotaliReferenti": helper.PostiPrenotatiReferente(prenotazioni),
			"biglietti":            biglietti,
			"bigliettiCalcolo":     bigliettiPerTipo,
			"totaleBiglietti":      totaleBiglietti,
			"totaleCosti":          totaleCosti,
		}
		helper.BigliettiTotali(prenotazioni)
		c.HTML(http.StatusOK, "Prenotazioni.html", data)

	} else {

		var prenotazioni []models.Prenotazione
		initializers.DB.Where("data::date = ?", dateQuery).Order(c.Query("orderby")).Find(&prenotazioni)
		// Calcolo dei Biglietti
		bigliettiPerTipo = helper.BigliettiTotali(prenotazioni)

		var totaleBiglietti, totaleCosti int
		for _, val := range bigliettiPerTipo {
			totaleBiglietti += val.NumBiglietti
			totaleCosti += val.CostoTotale
		}

		data := gin.H{
			"prenotazioni":         prenotazioni,
			"spettacolo":           spettacolo,
			"date":                 date,
			"dateQuery":            dateQuery,
			"numeriPosti":          numeriPosti,
			"postiprenotati":       helper.GetPostiPrenotati(dateQuery),
			"listaAttori":          listaAttori,
			"postiTotaliPrenotati": helper.PostiTotaliPrenotati(prenotazioni),
			"postiTotaliReferenti": helper.PostiPrenotatiReferente(prenotazioni),
			"biglietti":            biglietti,
			"bigliettiCalcolo":     bigliettiPerTipo,
			"totaleBiglietti":      totaleBiglietti,
			"totaleCosti":          totaleCosti,
		}
		helper.BigliettiTotali(prenotazioni)
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

	// Biglietto
	prenotazione.BigliettoID = helper.ReturnIntFromString(c, c.PostForm("bigliettoID"))
	var biglietto models.Biglietto
	if err := initializers.DB.First(&biglietto, prenotazione.BigliettoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Biglietti non trovati"})
	}
	prenotazione.BigliettoTipo = biglietto.Tipo

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

	// Get Lista attori
	var listaAttori []models.Attore
	if err := initializers.DB.Find(&listaAttori).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attori non trovati"})
	}

	// Mi ricavo le date disponibili
	date := helper.DateRange(spettacolo.Inizio, spettacolo.Fine)

	// Get Biglietti
	var biglietti []models.Biglietto
	if err := initializers.DB.Find(&biglietti).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Biglietti non trovati"})
	}

	data := gin.H{
		"prenotazione": prenotazione,
		"numeriPosti":  slicePosti,
		"date":         date,
		"spettacolo":   spettacolo,
		"listaAttori":  listaAttori,
		"biglietti":    biglietti,
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

	// Biglietto
	prenotazione.BigliettoID = helper.ReturnIntFromString(c, c.PostForm("bigliettoID"))
	var biglietti models.Biglietto
	if err := initializers.DB.First(&biglietti, prenotazione.BigliettoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Biglietti non trovati"})
	}
	prenotazione.BigliettoTipo = biglietti.Tipo
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

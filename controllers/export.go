package controllers

import (
	"cdbp4/initializers"
	"cdbp4/models"
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EXCEL = https://zetcode.com/golang/excel/?utm_content=cmp-true

func ExportToCSVSpettacoli(c *gin.Context) {

	// Set our headers so browser will download the file
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=Spettacoli.csv")
	//Get all users
	
	var spettacolo []models.Spettacolo
	initializers.DB.Order(c.DefaultQuery("orderby", "ID")).Find(&spettacolo)
	// Create a CSV writer using our HTTP response writer as our io.Writer
	wr := csv.NewWriter(c.Writer)
	//Expected data format : Here the Solution
	var data [][]string
	data = append(data, []string{
		"ID",
		"Nome",
		"Data Inizio",
		"Data Fine",
		"Teatro",
	})
	for _, record := range spettacolo {
		row := []string{
			strconv.FormatUint(uint64(record.ID), 10),
			record.Nome,
			record.Inizio.Format("2006-01-02 15:04:05"),
			record.Fine.Format("2006-01-02 15:04:05"),
			record.TeatroName,
		}
		data = append(data, row)
	}

	if err := wr.WriteAll(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate CSV file",
		})
		return
	}
}

func ExportToCSVPrenotazioni(c *gin.Context) {
	// Ricavo il nome dello spettacolo
	id := c.Param("id")
	var spettacolo models.Spettacolo
	initializers.DB.First(&spettacolo, id)
	nomeSpettacolo := spettacolo.Nome
	// Prendo la Query dell'url se non c'è la valorizzo NODATA
	dateQuery := c.DefaultQuery("data", "NODATA")
	var prenotazione []models.Prenotazione
	if dateQuery == "NODATA" {
		//Get all users
		initializers.DB.Where("spettacolo_id = ?", id).Find(&prenotazione)
		// Set our headers so browser will download the file
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment;filename=Prenotazioni-"+nomeSpettacolo+".csv")
	} else {
		//Get all users
		initializers.DB.Where("data::date = ?", dateQuery).Find(&prenotazione)
		// Set our headers so browser will download the file
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment;filename=Prenotazioni-"+nomeSpettacolo+"-"+dateQuery+".csv")
	}

	// Create a CSV writer using our HTTP response writer as our io.Writer
	wr := csv.NewWriter(c.Writer)
	//Expected data format : Here the Solution
	var data [][]string
	data = append(data, []string{
		"ID",
		"Nome",
		"Posti",
		"Data",
		"Spettacolo",
	})
	for _, record := range prenotazione {
		row := []string{
			strconv.FormatUint(uint64(record.ID), 10),
			record.Nome,
			strconv.Itoa(record.Posti),
			record.Data.Format("2006-01-02 15:04:05"),
			record.SpettacoloName,
		}
		data = append(data, row)
	}

	if err := wr.WriteAll(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate CSV file",
		})
		return
	}
}

func ExportToCSVTeatri(c *gin.Context) {

	// Set our headers so browser will download the file
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=Teatri.csv")
	//Get all users
	var teatro []models.Teatro
	initializers.DB.Order(c.DefaultQuery("orderby", "ID")).Find(&teatro)
	// Create a CSV writer using our HTTP response writer as our io.Writer
	wr := csv.NewWriter(c.Writer)
	//Expected data format : Here the Solution
	var data [][]string
	data = append(data, []string{
		"ID",
		"Nome",
		"Posti",
	})
	for _, record := range teatro {
		row := []string{
			strconv.FormatUint(uint64(record.ID), 10),
			record.Nome,
			strconv.Itoa(record.Posti),
		}
		data = append(data, row)
	}

	if err := wr.WriteAll(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate CSV file",
		})
		return
	}
}

func ExportToCSVAttori(c *gin.Context) {

	// Set our headers so browser will download the file
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=Attori.csv")
	//Get all users
	var attori []models.Attore
	initializers.DB.Order(c.DefaultQuery("orderby", "ID")).Find(&attori)
	// Create a CSV writer using our HTTP response writer as our io.Writer
	wr := csv.NewWriter(c.Writer)
	//Expected data format : Here the Solution
	var data [][]string
	data = append(data, []string{
		"ID",
		"Nome",
		"Cognome",
		"Email",
	})
	for _, record := range attori {
		row := []string{
			strconv.FormatUint(uint64(record.ID), 10),
			record.Nome,
			record.Cognome,
			record.Email,
		}
		data = append(data, row)
	}

	if err := wr.WriteAll(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate CSV file",
		})
		return
	}
}

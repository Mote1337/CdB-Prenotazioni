package helper

import (
	"cdbp4/initializers"
	"cdbp4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPostiPrenotati(dateQuery string) int {
	var postiprenotati models.Prenotazione
	if err := initializers.DB.Table("prenotaziones").Where("data::date = ? AND deleted_at IS NULL", dateQuery).Select("sum(posti) as postiprenotati").Scan(&postiprenotati.Posti).Error; err != nil {
		return 0
	}

	return postiprenotati.Posti
}

// GetSpettacoloByID Riceve un gin.Context in ingresso e ricavandosi l'id dello spettacolo ritorna la entry singola dello spettacolo con stesso id
func GetSpettacoloByID(c *gin.Context) models.Spettacolo {
	// spettacolo Variabile da popolare con il risultato della query sul DB
	var spettacolo models.Spettacolo
	// id Parametro dell'url del gin.Context che rispecchia l'id dello spettacolo in gorm.Model
	id := c.Param("id")
	// Assegnazione risultato query sulla variabile spettacolo
	if err := initializers.DB.First(&spettacolo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Spettacolo non trovato"})
	}
	// Ritorno del models.Spettacolo
	return spettacolo
}

// ReturnIntFromString riceve una String e la converte in Int
func ReturnIntFromString(c *gin.Context, number string) int {
	intero, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid String value to convert in Int"})
	}

	return int(intero)
}

// ReturnDatefromString
/*
func ReturnDatefromString(c *gin.Context, dataStringa string) time.Time {
	data, err := time.Parse("2006-01-02", dataStringa)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid String value to convert in time.Time"})
	}

	return data
}
*/

// ReturnUIntFromString riceve una String e la converte in Int
func ReturnUIntFromString(c *gin.Context, number string) uint {
	intero, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid String value to convert in UInt"})
	}

	return uint(intero)
}

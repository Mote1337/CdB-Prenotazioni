package helper

import "cdbp4/models"

func PostiTotaliPrenotati(prenotazioni []models.Prenotazione) int {
	var totale int
	for _, prenotazione := range prenotazioni {
		totale += prenotazione.Posti
	}
	return totale
}

package helper

import (
	"cdbp4/models"
)

func PostiTotaliPrenotati(prenotazioni []models.Prenotazione) int {
	var totale int
	for _, prenotazione := range prenotazioni {
		totale += prenotazione.Posti
	}
	return totale
}

// GPT
func PostiPrenotatiReferente(prenotazioni []models.Prenotazione) []models.Posti {

	// Mappa per rimuovere i duplicati e ottenere la lista dei nomi unici
	uniqueNomi := make(map[string]bool)
	for _, prenotazione := range prenotazioni {
		uniqueNomi[prenotazione.Referente] = true
	}

	// Lista dei nomi unici
	nomi := []string{}
	for nome := range uniqueNomi {
		nomi = append(nomi, nome)
	}

	// Lista dei posti per ogni nome
	postiPerNome := make(map[string]int)
	for _, nome := range nomi {
		for _, prenotazione := range prenotazioni {
			if prenotazione.Referente == nome {
				postiPerNome[nome] += prenotazione.Posti
			}
		}
	}

	// Creazione della lista di struct Posti
	var posti []models.Posti
	for _, nome := range nomi {
		p := models.Posti{
			Nome:  nome,
			Posti: postiPerNome[nome],
		}
		posti = append(posti, p)
	}

	return posti
}

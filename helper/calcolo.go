package helper

import (
	"cdbp4/initializers"
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

/*
func BigliettiTotali(prenotazioni []models.Prenotazione) {

	//Mi prendo lo struct di biglietti
	var biglietti []models.Biglietto
	if err := initializers.DB.Find(&biglietti).Error; err != nil {
		panic(err)
	}

	bigliettiPerBiglietto := make(map[string]int)
	for _, biglietto := range biglietti {
		for _, prenotazione := range prenotazioni {
			if biglietto.Tipo == prenotazione.BigliettoTipo {
				// Devo sommare i numeri dei biglietti
				// Devo sommare il costo dei biglietti
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
*/
// GPT
func BigliettiTotali(prenotazioni []models.Prenotazione) map[string]struct {
	NumBiglietti int
	CostoTotale  int
} {
	//Mi prendo lo struct di biglietti
	var biglietti []models.Biglietto
	if err := initializers.DB.Find(&biglietti).Error; err != nil {
		panic(err)
	}

	// Map per memorizzare il numero di biglietti e il costo totale per ogni tipo di biglietto
	bigliettiPerTipo := make(map[string]struct {
		NumBiglietti int
		CostoTotale  int
	})

	for _, biglietto := range biglietti {
		for _, prenotazione := range prenotazioni {
			if biglietto.Tipo == prenotazione.BigliettoTipo {
				// Aggiorna il numero di biglietti e il costo totale per questo tipo di biglietto
				if val, ok := bigliettiPerTipo[biglietto.Tipo]; ok {
					val.NumBiglietti += prenotazione.Posti
					val.CostoTotale += biglietto.Costo * prenotazione.Posti
					bigliettiPerTipo[biglietto.Tipo] = val
				} else {
					bigliettiPerTipo[biglietto.Tipo] = struct {
						NumBiglietti int
						CostoTotale  int
					}{
						NumBiglietti: prenotazione.Posti,
						CostoTotale:  biglietto.Costo * prenotazione.Posti,
					}
				}
			}
		}
	}

	// Stampa i risultati
	/*
		for tipo, val := range bigliettiPerTipo {
			fmt.Printf("Tipo di biglietto: %s - Numero di biglietti: %d - Costo totale: %d\n", tipo, val.NumBiglietti, val.CostoTotale)
		}
	*/
	return bigliettiPerTipo
}

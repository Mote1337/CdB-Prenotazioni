package models

import (
	"time"

	"gorm.io/gorm"
)

type Spettacolo struct {
	gorm.Model
	Nome       string
	Inizio     time.Time
	Fine       time.Time
	TeatroID   uint
	TeatroName string
}

type Prenotazione struct {
	gorm.Model
	Nome           string
	Posti          int
	BigliettoTipo  string
	BigliettoID    int
	Data           time.Time
	SpettacoloID   uint
	SpettacoloName string
	Referente      string
}

type Teatro struct {
	gorm.Model
	Nome  string
	Posti int
}

type Attore struct {
	gorm.Model
	Nome    string
	Cognome string
	Email   string
}

type Biglietto struct {
	gorm.Model
	Tipo  string
	Costo int
}

type Posti struct {
	Nome  string
	Posti int
}

type Ticket struct {
	gorm.Model
	Name  string
	Price Money
}

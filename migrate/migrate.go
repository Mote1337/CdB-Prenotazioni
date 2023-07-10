package main

import (
	"cdbp4/initializers"
	"cdbp4/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Spettacolo{}, &models.Prenotazione{}, &models.Teatro{})
}

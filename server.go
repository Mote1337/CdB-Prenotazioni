package main

import (
	"cdbp4/controllers"
	"cdbp4/initializers"
	"cdbp4/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Prenotazione{}, &models.Spettacolo{}, &models.Teatro{}, &models.Attore{})
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	spettacoliRoutes := r.Group("/spettacoli")
	{
		// CRUD - CREATE
		spettacoliRoutes.POST("/create", controllers.CreateSpettacolo)
		// CRUD - READ
		spettacoliRoutes.GET("/read", controllers.ReadSpettacoli)
		// CRUD - UPDATE
		spettacoliRoutes.GET("/update/:id", controllers.GetUpdateSpettacolo)
		spettacoliRoutes.POST("update/:id", controllers.PostUpdateSpettacolo)
		// CRUD - DELETE
		spettacoliRoutes.POST("/delete/:id", controllers.DeleteSpettacolo)
	}

	prenotazioniRoutes := r.Group("/prenotazioni")
	{
		// CRUD - CREATE
		prenotazioniRoutes.POST("/create", controllers.PostCreatePrenotazione)
		// CRUD - READ
		prenotazioniRoutes.GET("/read/:id", controllers.ReadPrenotazioniUnico)
		// CRUD - UPDATE
		prenotazioniRoutes.GET("/update/:id", controllers.GetUpdatePrenotazione)
		prenotazioniRoutes.POST("/update/:id", controllers.PostUpdatePrenotazione)
		// CRUD - DELETE
		prenotazioniRoutes.POST("/delete/:id", controllers.DeletePrenotazione)
	}

	teatriRoutes := r.Group("/teatri")
	{
		// CRUD - CREATE
		teatriRoutes.POST("/create", controllers.CreateTeatri)
		// CRUD - READ
		teatriRoutes.GET("/read", controllers.ReadTeatri)
		// CRUD - UPDATE
		teatriRoutes.GET("/update/:id", controllers.GetUpdateTeatri)
		teatriRoutes.POST("/update/:id", controllers.UpdateTeatri)
		// CRUD - DELETE
		teatriRoutes.POST("/delete/:id", controllers.DeleteTeatri)
	}

	attoriRoutes := r.Group("/attori")
	{
		// CRUD - CREATE
		attoriRoutes.POST("/create", controllers.CreateAttore)
		// CRUD - READ
		attoriRoutes.GET("/read", controllers.ReadAttori)
		// CRUD - UPDATE
		attoriRoutes.GET("/update/:id", controllers.GetUpdateAttore)
		attoriRoutes.POST("/update/:id", controllers.UpdateAttore)
		// CRUD - DELETE
		attoriRoutes.POST("/delete/:id", controllers.DeleteAttore)
	}

	exportRoutes := r.Group("/export")
	{
		exportRoutes.GET("csv/spettacoli", controllers.ExportToCSVSpettacoli)
		exportRoutes.GET("csv/prenotazioni/:id", controllers.ExportToCSVPrenotazioni)
		exportRoutes.GET("csv/teatri", controllers.ExportToCSVTeatri)
		exportRoutes.GET("csv/attori", controllers.ExportToCSVAttori)
	}
	r.GET("/", controllers.ReadRoot)
	r.Run(":3000")
}

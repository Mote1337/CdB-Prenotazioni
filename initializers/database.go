package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=rogue.db.elephantsql.com user=yfyqbqiy password=3xyppyNlyrsxvUvNh8s41UnTPecdzsD4 dbname=yfyqbqiy port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect DB")
	}
}

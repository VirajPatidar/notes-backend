package database

import (
	"github.com/VirajPatidar/NotesBackend/models"
	"gorm.io/driver/postgres"
    _ "github.com/heroku/x/hmetrics/onload"
    _ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
	"os"
	"fmt"
)

var DB *gorm.DB

func Connect() {

	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")
	port := os.Getenv("PGPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, pass, dbname, port) //Build connection string
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.Debug().AutoMigrate(&models.User{}, &models.Note{})
}

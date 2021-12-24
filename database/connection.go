package database

import (
	"github.com/VirajPatidar/NotesBackend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"fmt"
)

var DB *gorm.DB

func Connect() {

	host := os.Getenv("Host")
	user := os.Getenv("User")
	pass := os.Getenv("Password")
	dbname := os.Getenv("Database")
	port := os.Getenv("Port")

	// https://github.com/go-gorm/postgres
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, pass, port) //Build connection string
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Note{})
}

package database

import (
	"github.com/VirajPatidar/NotesBackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("Note.db"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

}

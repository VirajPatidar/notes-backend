package models
import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Notes []Note `json:"-"`
}

type Note struct {
	gorm.Model
	Title string
	Category string
	Details string
	UserID uint
}
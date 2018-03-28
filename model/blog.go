package model

import (
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Blog struct blog
type Blog struct {
	gorm.Model
	Title   string
	Content string
}

// Table name
const table = "blogs"

func (b Blog) Validation() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Title, validation.Required, validation.Length(5, 100)),
		validation.Field(&b.Content, validation.Required),
	)
}

func (b Blog) TableName() string {
	dbName := os.Getenv("DB_DATABASE")
	return dbName + "." + table
}

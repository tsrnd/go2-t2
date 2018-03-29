package model

import (
	"os"

	"github.com/jinzhu/gorm"
)

// Blog struct blog
type Blog struct {
	gorm.Model
	Title   string `validate:"required"`
	Content string `validate:"required"`
}

// Table name
const table = "blogs"

func (b Blog) TableName() string {
	dbName := os.Getenv("DB_DATABASE")
	return dbName + "." + table
}

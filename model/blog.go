package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Blog struct blog
type Blog struct {
	gorm.Model
	Title   string
	Content string
}

func (b Blog) Validation(title string, content string) error {
	return validation.Errors{
		"title":   validation.Validate(title, validation.Required, validation.Length(1, 100)),
		"content": validation.Validate(content, validation.Required),
	}.Filter()
}

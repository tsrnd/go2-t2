package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Blog struct blog
type Blog struct {
	Model
	Title   string
	Content string
}

func (b Blog) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Title, validation.Required, validation.Length(5, 50)),
		validation.Field(&b.Content, validation.Required),
	)
}

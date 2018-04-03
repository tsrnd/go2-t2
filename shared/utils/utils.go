package utils

import (
	"encoding/hex"
	"reflect"

	"golang.org/x/crypto/scrypt"
)

// GetStructTag get struct tag.
func GetStructTag(typeData interface{}, tagName string) (string, error) {
	field, ok := reflect.TypeOf(typeData).Elem().FieldByName(tagName)
	if !ok {
		return "", ErrorsNew("can't get Tag")
	}
	return string(field.Tag), nil
}

// CreateHashFromPassword scrypt hash password.
func CreateHashFromPassword(salt, password string) string {
	converted, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 16)
	return hex.EncodeToString(converted[:])
}

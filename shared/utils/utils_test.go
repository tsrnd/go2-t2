package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type gettag struct {
	a string `a:"1"`
}

func TestGetTagOK(t *testing.T) {
	testtag := &gettag{a: "test"}
	data, err := GetStructTag(testtag, "a")
	assert.Equal(t, "a:\"1\"", data)
	assert.NoError(t, err)
}

func TestGetTagNG(t *testing.T) {
	testtag := &gettag{a: "test"}
	data, err := GetStructTag(testtag, "b")
	assert.Equal(t, "", data)
	assert.Error(t, err)
}

func TestCreateHashFromPasswordOK(t *testing.T) {
	password := CreateHashFromPassword("test", "123456")
	assert.NotEmpty(t, password)
	assert.NotEqual(t, "123456", password)
}

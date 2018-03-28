package model

import "github.com/jinzhu/gorm"

//DBCon dbcon
var DBCon *gorm.DB

type Blog struct {
	ID      int
	Title   string
	Content string
}

//SetDatabase return DBCon
func SetDatabase(database *gorm.DB) {
	DBCon = database
}

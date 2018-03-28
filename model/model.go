package model

import (
	"github.com/jinzhu/gorm"
)

// Model general struct
// type Model struct {
// 	gorm.BaseModel
// }

//DBCon dbcon
var DBCon *gorm.DB

//SetDatabase return DBCon
func SetDatabase(database *gorm.DB) {
	DBCon = database
}

package model

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

// Model general struct
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//DBCon dbcon
var DBCon *gorm.DB

//SetDatabase return DBCon
func SetDatabase(database *gorm.DB) {
	DBCon = database
}

//Redirect redirect
func Redirect(w http.ResponseWriter, r *http.Request, target string) {
	// remove/add not default ports from req.Host
	http.Redirect(w, r, target, 301)
}

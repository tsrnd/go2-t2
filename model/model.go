package model

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

//DBCon dbcon
var DBCon *gorm.DB
var err error

// Blog type struct
type Blog struct {
	ID      int
	Title   string
	Content string
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Model general struct
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//SetDatabase return DBCon
func SetDatabase(database *gorm.DB) {
	DBCon = database
}

// GetAllPosts return all posts
func GetAllPosts() []Blog {
	var blogs []Blog
	DBCon.Select("id, title, content").Order("id desc, title").Find(&blogs)
	return blogs
}

//Redirect redirect
func Redirect(w http.ResponseWriter, r *http.Request) {
	// remove/add not default ports from req.Host
	target := "http://" + r.Host
	http.Redirect(w, r, target, 301)
}

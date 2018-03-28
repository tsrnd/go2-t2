package model

import (
	"github.com/jinzhu/gorm"
)

//DBCon dbcon
var DBCon *gorm.DB

// Blog is type of blog
type Blog struct {
	ID      int
	Title   string
	Content string
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

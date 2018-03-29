package repository

import (
	model "go2-t2/model"
	"log"
)

// GetAllPosts return all posts
func GetAllPosts() []model.Blog {
	var blogs []model.Blog
	allPosts := model.DBCon.Select("id, title, content").Order("id desc, title").Find(&blogs)
	if allPosts.Error != nil {
		log.Fatalln(allPosts.Error)
	}
	return blogs
}

package repository

import (
	model "go2-t2/model"
)

// GetAllPosts return all posts
func GetAllPosts() []model.Blog {
	var blogs []model.Blog
	model.DBCon.Select("id, title, content").Order("id desc, title").Find(&blogs)
	return blogs
}

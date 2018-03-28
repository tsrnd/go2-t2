package repository

import (
	model "go2-t2/model"
)

// GetAllPosts return all posts
func GetAllPosts() []model.Blog {
	var blogs []model.Blog
	model.DBCon.Find(&blogs)
	return blogs
}

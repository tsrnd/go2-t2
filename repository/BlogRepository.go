package repository

import (
	"go2-t2/model"
	"log"
)

func Store(blog model.Blog) {
	model.DBCon.NewRecord(blog)
	model.DBCon.Create(&blog)
}

func Get(id string) model.Blog {
	blog := model.Blog{}
	model.DBCon.First(&blog, id)
	return blog
}

func Update(id string) {
	blog := model.Blog{}
	model.DBCon.First(&blog, id)

	model.DBCon.Model(&blog).UpdateColumn("title", "content")
}

func GetAllPosts() []model.Blog {
	var blogs []model.Blog
	allPosts := model.DBCon.Select("id, title, content").Order("id desc, title").Find(&blogs)
	if allPosts.Error != nil {
		log.Fatalln(allPosts.Error)
	}
	return blogs
}

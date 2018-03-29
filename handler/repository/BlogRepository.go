package repository

import "go2-t2/model"

func Store(blog model.Blog) {
	model.DBCon.NewRecord(blog)
	model.DBCon.Create(&blog)
}

package repository

import (
	"go2-t2/model"
)

// Delete will delete post given id
func Delete(id string) {
	blog := model.Blog{}
	model.DBCon.First(&blog, id)

	model.DBCon.Delete(&blog)
}

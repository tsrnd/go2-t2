package user

import (
	"github.com/tsrnd/trainning/shared/gorm/model"
)

// ----------------------------------------------------------
// Database
// ----------------------------------------------------------

// User table struct.
// http://jinzhu.me/gorm/models.html#model-definition
type User struct {
	ID       uint64 `gorm:"column:id_user_app;primary_key"`
	UUID     string `gorm:"column:uuid;type:char(36)"`
	UserName string `gorm:"column:user_name;type:varchar(20)"`
	model.BaseModel
}

// TableName function custom table name.
func (User) TableName() string {
	return "user_app"
}

// GetCustomClaims get customs claims
func (u User) GetCustomClaims() map[string]interface{} {
	claims := make(map[string]interface{})
	userclaim := struct {
		ID uint64 `json:"id"`
	}{
		ID: u.ID,
	}
	claims["user"] = userclaim
	return claims
}

// GetIdentifier get identifier function
func (u User) GetIdentifier() uint64 {
	return u.ID
}

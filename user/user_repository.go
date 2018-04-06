package user

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/utils"
)

// RepositoryInterface interface.
type RepositoryInterface interface {
	FindOrCreate(string) (User, error)
	Delete(uint64) (int64, error)
	// FindUserByID(uint64) (User, error)
}

// Repository struct.
type Repository struct {
	repository.BaseRepository
	// connect master database.
	masterDB *gorm.DB
	// connect read replica database.
	readDB *gorm.DB
	// redis connect Redis.
	redis *redis.Conn
}

// // FindUserID find User by ID.
// func (r *Repository) FindUserByID(userID uint64) (User, error) {
// 	user := User{}
// 	err := r.masterDB.Where("id_user_app = ?", userID).First(&user).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return user, err
// 	}
// 	return user, utils.ErrorsWrap(err, "can't find")
// }

// Delete find user.
func (r *Repository) Delete(id uint64) (int64, error) {
	stm := r.masterDB.Where("id_user_app = ?", id).Delete(&User{})
	if stm.Error != nil {
		return stm.RowsAffected, utils.ErrorsWrap(stm.Error, "Can't not delete")
	}
	return stm.RowsAffected, nil
}

// FindOrCreate find user by uuid or create if uuid is not existed in DB.
func (r *Repository) FindOrCreate(uuid string) (User, error) {
	user := User{UUID: uuid}
	err := r.readDB.FirstOrCreate(&user, user).Error
	return user, utils.ErrorsWrap(err, "Can't first or create")
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB, redis *redis.Conn) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read, redis: redis}
}

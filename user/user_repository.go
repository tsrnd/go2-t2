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
	CreateUser(string, string, *gorm.DB) (uint64, error)
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

// FindOrCreate find user by uuid or create if uuid is not existed in DB.
func (r *Repository) FindOrCreate(uuid string) (User, error) {
	user := User{UUID: uuid}
	err := r.masterDB.FirstOrCreate(&user, user).Error
	return user, utils.ErrorsWrap(err, "Can't first or create")
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB, redis *redis.Conn) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read, redis: redis}
}

// CreateUser create user
func (r *Repository) CreateUser(uuid string, username string, tx *gorm.DB) (uint64, error) {
	user := User{UUID: uuid, UserName: username}
	result := tx.Create(&user)
	return user.ID, utils.ErrorsWrap(result.Error, "can create user")
}

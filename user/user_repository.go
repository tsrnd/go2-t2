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
	CreateUserApp(string, string) (uint64, error)
	FindAll() ([]User, error)
	First(uint64) (User, error)
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

// FindAll find all users
func (r *Repository) FindAll() ([]User, error) {
	users := []User{}
	err := r.readDB.Find(&users).Error
	if err != nil {
		return users, utils.ErrorsWrap(err, "can't get user")
	}

	return users, nil
}

// First user by id in DB.
func (r *Repository) First(id uint64) (User, error) {
	user := User{ID: id}
	err := r.readDB.First(&user).Error
	return user, utils.ErrorsWrap(err, "Can't find user")
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB, redis *redis.Conn) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read, redis: redis}
}

// CreateUserApp create user app
func (r *Repository) CreateUserApp(uuid string, username string) (uint64, error) {
	user := User{UUID: uuid, UserName: username}
	result := r.masterDB.Create(&user)

	return user.ID, utils.ErrorsWrap(result.Error, "Can't create user app")
}

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
	FindAll() ([]User, error)
	First(uint64) (User, error)
	Update(uint64, string) error
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

// Update user with user_name
func (r *Repository) Update(idUserApp uint64, userName string) error {
	user := User{}
	r.masterDB.First(&user, idUserApp)
	err := r.masterDB.Model(&user).Updates(User{UserName: userName}).Error
	return utils.ErrorsWrap(err, "Can't update")
}

// FindAll find all users
func (r *Repository) FindAll() ([]User, error) {
	users := []User{}
	err := r.readDB.Find(&users).Error
	return users, err
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

package user

import (
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/authentication"
	"github.com/tsrnd/trainning/shared/usecase"
	"github.com/tsrnd/trainning/shared/utils"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
	RegisterByDevice(uuid string) (PostRegisterByDeviceResponse, error)
	Register(request PostRegisterRequest) error
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
}

// RegisterByDevice func.
func (u *Usecase) RegisterByDevice(uuid string) (response PostRegisterByDeviceResponse, err error) {
	// var userID int64
	response = PostRegisterByDeviceResponse{}
	user, err := u.repository.FindOrCreate(uuid)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repositoryInterface.FindOrCreate() error")
	}
	// store user to JWT
	response.Token, err = authentication.GenerateToken(user)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repository.GenerateJWToken() error")
	}
	return
}

// Register func.
func (u *Usecase) Register(request PostRegisterRequest) error {
	// var userID int64
	tx := u.db.Begin()
	_, err := u.repository.CreateUser(request.UUID, request.UserName, tx)

	if err != nil {
		return utils.ErrorsWrap(err, "cant create user")
	}
	tx.Commit()
	return nil
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}

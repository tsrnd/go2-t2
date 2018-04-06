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
	DeleteUser(userID uint64) (DeleteUserResponse, error)
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

// Delete user func.
func (u *Usecase) DeleteUser(userID uint64) (DeleteUserResponse, error) {
	response := DeleteUserResponse{}
	rowsAffected, err := u.repository.Delete(userID)
	response.Message = "Delete success"
	if err != nil {
		response.Message = "Error when delete user"
		return response, err
	}
	if rowsAffected == 0 {
		response.Message = "Cant find user"
	}

	return response, nil
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}

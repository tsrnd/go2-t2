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
	UpdateUser(PutUpdateByUserRequest) (CommonResponse, error)
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

// UpdateUser func
func (u *Usecase) UpdateUser(request PutUpdateByUserRequest) (CommonResponse, error) {
	response := CommonResponse{Message: "Update success", Errors: nil}
	err := u.repository.Update(request.ID, request.UserName)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repositoryInterface.Update() error")
	}
	return response, err
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}

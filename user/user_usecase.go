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
	Create(request PostCreateRequest) (PostCreateResponse, error)
	GetAllUsers() ([]GetUserResponse, error)
	GetUserByID(id uint64) (GetGetUserByIDResponse, error)
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

// Create func.
func (u *Usecase) Create(request PostCreateRequest) (response PostCreateResponse, err error) {
	response = PostCreateResponse{}
	response.ID, err = u.repository.CreateUserApp(request.UUID, request.UserName)

	return response, utils.ErrorsWrap(err, "repository.CreateUser() error")
}

// GetAllUsers func.
func (u *Usecase) GetAllUsers() (response []GetUserResponse, err error) {
	users, err := u.repository.FindAll()

	response = []GetUserResponse{}
	for _, user := range users {
		response = append(response, GetUserResponse{
			ID:       user.ID,
			UUID:     user.UUID,
			UserName: user.UserName,
		})
	}

	return
}

// GetUserByID func.
func (u *Usecase) GetUserByID(id uint64) (response GetGetUserByIDResponse, err error) {
	user, err := u.repository.First(id)

	if err != nil {
		return response, utils.ErrorsWrap(err, "repositoryInterface.First() error")
	}

	response = GetGetUserByIDResponse{
		ID:       user.ID,
		UUID:     user.UUID,
		UserName: user.UserName,
	}

	return
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}

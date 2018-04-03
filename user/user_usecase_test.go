package user

import (
	"errors"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tsrnd/trainning/shared/test"
	"github.com/tsrnd/trainning/shared/usecase"
)

func TestPostRegisterByDeviceSuccess(t *testing.T) {
	user := User{ID: 100, UUID: uuid}
	uc := InitializeUsecaseWith(nil, user)

	response, err := uc.RegisterByDevice(uuid)
	arrayString := strings.Split(response.Token, ".")

	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.NotEqual(t, "", response.Token)
	assert.Equal(t, 3, len(arrayString))
}

func TestPostRegisterByDeviceFailWhenFirstOrCreateHasError(t *testing.T) {
	user := User{UUID: uuid}
	uc := InitializeUsecaseWith(errors.New("Repository.FirstOrCreate has an error"), user)

	_, err := uc.RegisterByDevice(uuid)

	assert.Regexp(t, "(.+)Repository.FirstOrCreate has an error$", err.Error())
	assert.Error(t, err)
}

func TestPostRegisterByDeviceFailWhenGenerateTokenHasError(t *testing.T) {
	user := User{}
	uc := InitializeUsecaseWith(nil, user)

	_, err := uc.RegisterByDevice(uuid)

	assert.Regexp(t, "Object is empty$", err.Error())
	assert.Error(t, err)
}

// MocksUserRepo struct
type MocksUserRepo struct {
	User            User
	FindOrCreateErr error
}

// InitializeUsecaseWith error and response mock
func InitializeUsecaseWith(findOrCreateErr error, user User) *Usecase {
	_, gorm := test.NewSQLMock()
	var logger logrus.Logger
	faker.FakeData(&logger)
	bUC := usecase.BaseUsecase{Logger: &logger}
	mockRepo := MocksUserRepo{User: user, FindOrCreateErr: findOrCreateErr}
	return NewUsecase(&bUC, gorm, &mockRepo)
}

// FindOrCreate func
func (uc *MocksUserRepo) FindOrCreate(uuid string) (User, error) {
	return uc.User, uc.FindOrCreateErr
}

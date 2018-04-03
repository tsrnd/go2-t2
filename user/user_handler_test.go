package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/handler"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/test"
	"github.com/tsrnd/trainning/shared/usecase"
)

const (
	// success data
	validDeviceID  = "123e4567-e89b-12d3-a456-426655440002"
	validToken     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTY1OTU1MzEsImlhdCI6MTUxNjU5MTkzMSwiaXNzIjoiZnItY2lyY2xlLWFwaS5jb20iLCJuYmYiOjE1MTY1OTE5MzEsInVzZXJfaWQiOiJcdTAwMDcifQ.3EVeZKRTiZQgOY4wa0QfIleMO8mPBEA8Kl8EFrpVC5Y"
	validContent   = "application/x-www-form-urlencoded"
	invalidContent = "application/x-www-form-urlencoded, test"
)

type PostRegisterByDeviceParam struct {
	DeviceID string
}

type mockUsecase struct {
	UsecaseInterface
	FakeRegisterByDevice func(string) (PostRegisterByDeviceResponse, error)
}

func (u *mockUsecase) RegisterByDevice(DeviceID string) (PostRegisterByDeviceResponse, error) {
	return u.FakeRegisterByDevice(DeviceID)
}

func getSuccessResponsePostRegisterByDevice(token string) (PostRegisterByDeviceResponse, error) {
	var response PostRegisterByDeviceResponse
	response.Token = token
	return response, nil
}

func getErrorResponsePostRegisterByDevice(message string) (PostRegisterByDeviceResponse, error) {
	var response PostRegisterByDeviceResponse
	return response, errors.New(message)
}

func request(h *HTTPHandler, method string, contentType string, path string, param *PostRegisterByDeviceParam) (*httptest.ResponseRecorder, error) {
	form := url.Values{}
	form.Add("device_id", param.DeviceID)
	req, err := http.NewRequest(method, path, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	rec := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(h.RegisterByDevice)
	handlerFunc.ServeHTTP(rec, req)

	return rec, nil
}

func TestV1PostRegisterByDeviceSuccess(t *testing.T) {
	loggerConfig := test.NewLogger()
	mockUsecase := &mockUsecase{
		FakeRegisterByDevice: func(DeviceID string) (PostRegisterByDeviceResponse, error) {
			return getSuccessResponsePostRegisterByDevice(validToken)
		},
	}
	h := &HTTPHandler{
		*handler.NewBaseHTTPHandler(loggerConfig),
		mockUsecase,
	}

	tests := []struct {
		description string
		method      string
		path        string
		params      *PostRegisterByDeviceParam
	}{
		{
			"login by valid DeviceID",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: validDeviceID,
			},
		},
	}
	tt := tests[0]
	t.Run(tt.description, func(t *testing.T) {
		// request
		rec, err := request(h, tt.method, validContent, tt.path, tt.params)
		if err != nil {
			t.Errorf("error occured by precondition for test: %s", err)
			return
		}

		assert.Equal(t, http.StatusOK, rec.Code)

		// prepare expected response body
		var expected PostRegisterByDeviceResponse
		expected.Token = validToken

		var response PostRegisterByDeviceResponse
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, expected, response)
	})
}

var checkValidateDeviceIDFailedResponse = func(t *testing.T, method string, path string, params *PostRegisterByDeviceParam, h *HTTPHandler) {
	// request
	rec, err := request(h, method, validContent, path, params)
	if err != nil {
		t.Errorf("error occured by precondition for test: %s", err)
		return
	}

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var data CommonResponse
	json.Unmarshal(rec.Body.Bytes(), &data)
	assert.Equal(t, "validation error.", data.Message)
	assert.Equal(t, []string{"device_id is error."}, data.Errors)
}

var parseFormErrorBadRequestResponse = func(t *testing.T, method string, path string, params *PostRegisterByDeviceParam, h *HTTPHandler) {
	// request
	rec, err := request(h, method, invalidContent, path, params)
	if err != nil {
		t.Errorf("error occured by precondition for test: %s", err)
		return
	}

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var data CommonResponse
	json.Unmarshal(rec.Body.Bytes(), &data)
	assert.Equal(t, "Parse request error.", data.Message)
	assert.Empty(t, data.Errors)
}

var checkInternalErrorResponse = func(t *testing.T, method string, path string, params *PostRegisterByDeviceParam, h *HTTPHandler) {
	// request
	rec, err := request(h, method, validContent, path, params)
	if err != nil {
		t.Errorf("error occured by precondition for test: %s", err)
		return
	}

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestV1PostRegisterByDeviceFailed(t *testing.T) {
	loggerConfig := test.NewLogger()
	errorMessage := "something occured in usecase"

	mockUsecase := &mockUsecase{
		FakeRegisterByDevice: func(DeviceID string) (PostRegisterByDeviceResponse, error) {
			return getErrorResponsePostRegisterByDevice(errorMessage)
		},
	}
	h := &HTTPHandler{
		*handler.NewBaseHTTPHandler(loggerConfig),
		mockUsecase,
	}

	tests := []struct {
		description string
		method      string
		path        string
		params      *PostRegisterByDeviceParam
		wantFunc    func(*testing.T, string, string, *PostRegisterByDeviceParam, *HTTPHandler)
	}{
		{
			"ParseForm has error",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e4567-e89b-12d3-a456-426655440004",
			},
			parseFormErrorBadRequestResponse,
		},
		{
			"RegisterByDevice usecase fail",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e4567-e89b-12d3-a456-426655440004",
			},
			checkInternalErrorResponse,
		},
		{
			"DeviceID is empty",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"DeviceID has 36 character but not in format 8-4-4-4-12",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e4-567e89b-12d3-a4564266-55440002",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"DeviceID has length < 36 character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e4-567e89b-12d3-a4564266-5544",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"DeviceID has length > 36 characters",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e4-567e89b-12d3-a4564266-55440002dsadasd",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"DeviceID has special character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e4567-e89b-1$d3-a456-426655440004",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 1 of DeviceID has special character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e$567-e89b-12d3-a456-426655440002",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 1 of DeviceID has length > 8",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567s-e89b-12d3-a456-42665544002",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 1 of DeviceID has length < 8",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e356-es89b-12d3-a456-42665544002",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 2 of DeviceID has special character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e$9b-12d3-a456-426655440023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 2 of DeviceID has > 4 characters",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89ba-12d3-a456-42665540023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 2 of DeviceID has < 4 characters",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e9b-12d3-a456-426655s440023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 3 of DeviceID has special character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-1$d3-a456-426655440023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 3 of DeviceID has > 4 character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-152d3-a456-42655440023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 3 of DeviceID has < 4 characters",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-123-a456-4266554440023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 4 of DeviceID has special character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-12d3-a$56-426655440023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 4 of DeviceID has > 4 character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-12d3-aa456-46655440023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 4 of DeviceID has < 4 character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-12d3-a56-426655440023",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 5 of DeviceID has special character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-12d3-a456-426655440$23",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 5 of DeviceID has > 12 character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-12d3-a456-426655440023a",
			},
			checkValidateDeviceIDFailedResponse,
		},
		{
			"Block 5 of DeviceID has < 12 character",
			"POST",
			"/v1/register/device",
			&PostRegisterByDeviceParam{
				DeviceID: "123e3567-e89b-12d3-a456-42665544002",
			},
			checkValidateDeviceIDFailedResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			tt.wantFunc(t, tt.method, tt.path, tt.params, h)
		})
	}
}

func TestNewHandlerSuccess(t *testing.T) {
	loggerConfig := test.NewLogger()
	// new base
	bh := handler.NewBaseHTTPHandler(loggerConfig)
	bu := usecase.NewBaseUsecase(loggerConfig)
	br := repository.NewBaseRepository(loggerConfig)
	// sql & cache
	_, gormdb := test.NewSQLMock()
	sql := infrastructure.SQL{Master: gormdb, Read: gormdb}
	cache := infrastructure.Cache{}
	// new http handler
	h := NewHTTPHandler(bh, bu, br, &sql, &cache)
	assert.NotNil(t, h)
}

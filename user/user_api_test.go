package user

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/handler"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/test"
	"github.com/tsrnd/trainning/shared/usecase"
)

var (
	apiTest       *test.APITest
	circleBaseDir = os.Getenv("FR_CIRCLE_API_DIR")
)

func TestMain(m *testing.M) {
	var mux *chi.Mux
	mux, apiTest = test.NewTestMain(m, setupHandler)
	server := httptest.NewServer(mux)
	code := m.Run()

	defer infrastructure.CloseLogger(apiTest.TRouter.LoggerHandler.Logfile)
	defer infrastructure.CloseRedis(apiTest.TRouter.CacheHandler.Conn)
	defer server.Close()
	defer truncateAllTableRelation()
	os.Exit(code)
}

func newDeviceID() string {
	deviceID := make([]byte, 16)
	n, _ := io.ReadFull(rand.Reader, deviceID)
	if n != len(deviceID) {
		return ""
	}
	deviceID[8] = deviceID[8]&^0xc0 | 0x80
	deviceID[6] = deviceID[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", deviceID[0:4], deviceID[4:6], deviceID[6:8], deviceID[8:10], deviceID[10:])
}

func setupHandler(r *test.Router) {
	// error handler set.
	eh := handler.NewHTTPErrorHandler(r.LoggerHandler.Log)
	r.Mux.NotFound(eh.StatusNotFound)
	r.Mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	// profiler
	env := os.Getenv("ENV_API")
	if env == "development" {
		r.Mux.Mount("/debug", middleware.Profiler())
	}

	// base set.
	bh := handler.NewBaseHTTPHandler(r.LoggerHandler.Log)
	// base set.
	br := repository.NewBaseRepository(r.LoggerHandler.Log)
	// base set.
	bu := usecase.NewBaseUsecase(r.LoggerHandler.Log)

	// item set.
	ah := NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler)

	r.Mux.Route("/v1", func(cr chi.Router) {
		cr.Post("/register/device", ah.RegisterByDevice)
	})
}

// createRequestBody func
func createRequestBody(fields map[string][]string) (body *bytes.Buffer, writer *multipart.Writer) {
	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	for key, values := range fields {
		for _, val := range values {
			writer.WriteField(key, val)
		}
	}
	writer.Close()
	return body, writer
}

// sendRequest func
func sendRequest(fields map[string][]string, contentType string, apiTest *test.APITest) *httptest.ResponseRecorder {
	form := url.Values{}
	for key, values := range fields {
		for _, val := range values {
			form.Add(key, val)
		}
	}
	req, _ := http.NewRequest("POST", "/v1/register/device", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", contentType)
	resp := apiTest.ExecuteRequest(req)
	return resp
}

// checkStatusCodeBadRequest func
func checkStatusCodeResponse(t *testing.T, resp *httptest.ResponseRecorder, expectCode int) {
	assert.NotNil(t, resp)
	assert.Equal(t, expectCode, resp.Code)
	return
}

// checkJSONResponeError func
func checkJSONResponeError(t *testing.T, resp *httptest.ResponseRecorder) {
	response := PostRegisterByDeviceResponse{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NotEmpty(t, response.Errors)
	assert.NotEmpty(t, response.Message)
	return
}

// checkJSONResponeMessage func
func checkJSONResponeMessage(t *testing.T, resp *httptest.ResponseRecorder, message string) {
	response := PostRegisterByDeviceResponse{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, response.Message, message)
	return
}

// Test Register by device fail with invalid UUID
func TestRegisterByDeviceFailWithInvalidUUID(t *testing.T) {
	// initialize
	apiTest.T = t
	testCaseStatusError := []struct {
		name         string
		paramRequest map[string][]string
	}{
		{
			name:         "miss device id parameter",
			paramRequest: map[string][]string{},
		},
		{
			name: "device id is empty",
			paramRequest: map[string][]string{
				"device_id": {""},
			},
		},
		{
			name: "device id has length > 36",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a456-426655440018MORETHAN36"},
			},
		},
		{
			name: "device id has length < 36",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a456-LOWER36"},
			},
		},
		{
			name: "device id wrong format 8-4-4-4-12",
			paramRequest: map[string][]string{
				"device_id": {"124567-e89b-12d3-a43333356-WRONGFORMAT"},
			},
		},
		{
			name: "block 1 of DeviceID has special character",
			paramRequest: map[string][]string{
				"device_id": {"123$4567-e89b-12d3-a456-426655440018"},
			},
		},
		{
			name: "block 1 of DeviceID has > 8 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567MORE8-e89b-12d3-a456-426655440018"},
			},
		},
		{
			name: "block 1 of DeviceID has < 8 character",
			paramRequest: map[string][]string{
				"device_id": {"LOWER8-e89b-12d3-a456-426655440018"},
			},
		},
		{
			name: "block 2 of DeviceID has special character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e$9b-12d3-a456-426655440018"},
			},
		},
		{
			name: "block 2 of DeviceID has > 4 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e8MORE89b-12d3-a456-426655440018"},
			},
		},
		{
			name: "block 2 of DeviceID has < 4 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-eb-12d3-a456-426655440018"},
			},
		},
		{
			name: "block 3 of DeviceID has special character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12$3-a456-426655440018"},
			},
		},
		{
			name: "block 3 of DeviceID has > 4 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d356-a456-426655440018"},
			},
		},
		{
			name: "block 3 of DeviceID has < 4 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12-a456-426655440018"},
			},
		},
		{
			name: "block 4 of DeviceID has special character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a$56-426655440018"},
			},
		},
		{
			name: "block 4 of DeviceID has > 4 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a456sdasdas-426655440018"},
			},
		},
		{
			name: "block 4 of DeviceID has < 4 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a6-426655440018"},
			},
		},
		{
			name: "block 5 of DeviceID has special character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a456-42665544001$"},
			},
		},
		{
			name: "block 5 of DeviceID has > 12 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a456-426655440018MORETHAN12"},
			},
		},
		{
			name: "block 5 of DeviceID has < 12 character",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a456-LOWER12"},
			},
		},
	}

	for _, testCase := range testCaseStatusError {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			resp := sendRequest(testCase.paramRequest, "application/x-www-form-urlencoded", apiTest)
			// check status bad request.
			checkStatusCodeResponse(t, resp, http.StatusBadRequest)
			// check response data.
			checkJSONResponeError(t, resp)
			// check user is not created in user_app table
			assert.False(t, checkUserExisted(testCase.paramRequest["device_id"][0]))
		})
	}
}

// Test Register by device fail with invalid content-type
func TestRegisterByDeviceFailWithInvalidContentType(t *testing.T) {
	// initialize
	apiTest.T = t
	testCaseStatusError := []struct {
		name         string
		paramRequest map[string][]string
	}{
		{
			name: "invalid content type",
			paramRequest: map[string][]string{
				"device_id": {"123e4567-e89b-12d3-a456-123456789123"},
			},
		},
	}
	t.Run(testCaseStatusError[0].name, func(t *testing.T) {
		resp := sendRequest(testCaseStatusError[0].paramRequest, "application/x-www-form-urlencoded, test", apiTest)
		// check status bad request.
		checkStatusCodeResponse(t, resp, http.StatusBadRequest)
		// check response data.
		checkJSONResponeMessage(t, resp, "Parse request error.")
		// check user is not created in user_app table
		assert.False(t, checkUserExisted(testCaseStatusError[0].paramRequest["device_id"][0]))
	})
}

// status ok with valid device id.
func TestV1RegisterDeviceStatusOK(t *testing.T) {
	apiTest.T = t
	deviceID := newDeviceID()
	truncateAllTableRelation()
	testCaseStatusOK := []struct {
		name         string
		paramRequest map[string][]string
	}{
		{
			name: "deviceID is not existed",
			paramRequest: map[string][]string{
				"device_id": {deviceID},
			},
		},
		{
			name: "deviceID is existed",
			paramRequest: map[string][]string{
				"device_id": {deviceID},
			},
		},
	}

	for _, testCase := range testCaseStatusOK {
		t.Run(testCase.name, func(t *testing.T) {
			var userExisted User
			if checkUserExisted(testCase.paramRequest["device_id"][0]) {
				userExisted, _ = getUserByUUID(testCase.paramRequest["device_id"][0])
			}
			resp := sendRequest(testCase.paramRequest, "application/x-www-form-urlencoded", apiTest)
			// check status OK.
			checkStatusCodeResponse(t, resp, http.StatusOK)
			response := &PostRegisterByDeviceResponse{}
			json.Unmarshal(resp.Body.Bytes(), response)
			assert.Empty(t, response.Errors)
			assert.Empty(t, response.Message)
			assert.Equal(t, 3, len(strings.Split(response.Token, ".")))
			// check not changed in db if user existed
			if userExisted.ID != 0 {
				userAfterRegister, _ := getUserByUUID(testCase.paramRequest["device_id"][0])
				assert.Equal(t, userAfterRegister.ID, userExisted.ID)
			}
		})
	}
}

func truncateAllTableRelation() {
	db := apiTest.TRouter.SQLHandler.Master
	db.Exec("TRUNCATE user_app RESTART IDENTITY CASCADE;")
}

func checkUserExisted(uuid string) bool {
	user, err := getUserByUUID(uuid)
	if err != nil || user.ID == 0 {
		return false
	}
	return true
}

func getUserByUUID(uuid string) (User, error) {
	db := apiTest.TRouter.SQLHandler.Read
	user := User{}
	err := db.Order("id_user_app desc").First(&user, User{UUID: uuid}).Error
	return user, err
}

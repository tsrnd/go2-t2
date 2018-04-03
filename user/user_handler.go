package user

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/handler"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/usecase"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.BaseHTTPHandler
	usecase UsecaseInterface
}

// RegisterByDevice to register user ID which originates from Device ID.
//
// "First": Search User from Entity by Device ID.
// "Second": If User record exists,move to step "Finally".
// "Third": If User record does not exist, register device ID to Entity.
// "Finally":store User_ID acquired from Entity to JSON Web Token (JWT).
func (h *HTTPHandler) RegisterByDevice(w http.ResponseWriter, r *http.Request) {
	// mapping post to struct.
	request := PostRegisterByDeviceRequest{}
	err := h.Parse(r, &request)
	if err != nil {
		common := CommonResponse{Message: "Parse request error.", Errors: nil}
		h.StatusBadRequest(w, common)
		return
	}

	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	// request login by uuid.
	response, err := h.usecase.RegisterByDevice(request.DeviceID)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecaseInterface.LoginByDevice() error")
		common := CommonResponse{Message: "Internal server error response.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// user set.
	userRepo := NewRepository(br, s.Master, s.Read, c.Conn)
	userUsecase := NewUsecase(bu, s.Master, userRepo)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: userUsecase}
}

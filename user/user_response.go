package user

// CommonResponse responses common json data.
type CommonResponse struct {
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

// PostRegisterByDeviceResponse response.
type PostRegisterByDeviceResponse struct {
	CommonResponse
	Token string `json:"token"`
}

// PostCreateRepository struct.
type PostCreateRepository struct {
	ID uint64 `json:"id"`
}

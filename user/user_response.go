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

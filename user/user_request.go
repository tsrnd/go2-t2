package user

// PostRegisterByDeviceRequest struct.
type PostRegisterByDeviceRequest struct {
	DeviceID string `form:"device_id" validate:"required,uuid"`
}

// PostRegisterRequest struct.
type PostRegisterRequest struct {
	UUID     string `form:"uuid" validate:"required"`
	UserName string `form:"user_name" validate:"required"`
}

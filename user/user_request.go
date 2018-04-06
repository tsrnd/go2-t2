package user

// PostRegisterByDeviceRequest struct.
type PostRegisterByDeviceRequest struct {
	DeviceID string `form:"device_id" validate:"required,uuid"`
}

//PutUpdateByUserRequest struct
type PutUpdateByUserRequest struct {
	ID       uint64
	UserName string `form:"username" validate:"required"`
}

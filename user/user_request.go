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
// PostCreateRequest struct.
type PostCreateRequest struct {
	UUID     string `form:"uuid" validate:"required"`
	UserName string `form:"user_name" validate:"required"`
}

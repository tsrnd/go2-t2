package user

// PostRegisterByDeviceRequest struct.
type PostRegisterByDeviceRequest struct {
	DeviceID string `form:"device_id" validate:"required,uuid"`
}

//DeleteUserRequest user struct
type DeleteUserRequest struct {
	IdUserApp uint64 `json:"id_user_app" validate:"required"`
}

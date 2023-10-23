package modelUser

type UpdateUserReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Position    string `json:"position" validate:"required"`
	Department  string `json:"department" validate:"required"`
}

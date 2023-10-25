package modelRoom

type CreateRoomReq struct {
	Room_Name string `json:"room_name" validate:"required"`
	Floor     int    `json:"floor" validate:"required"`
	Building  string `json:"building" validate:"required"`
}

type UpdateRoomReq struct {
	ID        int    `json:"id" validate:"required,gte=1"`
	Room_Name string `json:"room_name" validate:"required"`
	Floor     int    `json:"floor" validate:"required"`
	Building  string `json:"building" validate:"required"`
	Status    int    `json:"status" validate:"gte=0"`
}

package modelFeedback

type CreateFeedbackReq struct {
	Name_Feed_Back string   `json:"name_feed_back" validate:"required"`
	RoomID         uint     `json:"room_id" validate:"required,gte=1"`
	Description    string   `json:"description" validate:"required"`
	CategoryID     uint     `json:"category_id" validate:"required,gte=1"`
	UserID         uint     `json:"user_id" validate:"required,gte=1"`
	Urls           []string `json:"url" validate:"required"`
}

type RevicerFeedbackReq struct {
	Feedback_id uint `json:"feedback_id"`
	Option      int  `json:"option"`
	User_id     uint `json:"user_id"`
}

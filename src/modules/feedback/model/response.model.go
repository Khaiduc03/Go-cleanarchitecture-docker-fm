package modelFeedback

import "FM/src/entities"

type GetFeedbackRes struct {
	ID             uint     `json:"id"`
	Name_Feed_Back string   `json:"name_feed_back" validate:"required"`
	Room           entities.Room     `json:"room"`
	Description    string   `json:"description" validate:"required"`
	Category       entities.Category `json:"category"`
	User           entities.User     `json:"user"`
	Urls           []string `json:"url" validate:"required"`
}

type GetAllFeedbackRes struct {
	ID             uint   `json:"id"`
	Name_Feed_Back string `json:"name_feed_back" validate:"required"`
	RoomID         uint   `json:"room_id" validate:"required,gte=1"`
	Description    string `json:"description" validate:"required"`
	CategoryID     uint   `json:"category_id" validate:"required,gte=1"`
	UserID         uint   `json:"user_id" validate:"required,gte=1"`
}

type GetAllHistoryFeedbackRes struct {
	ID             uint              `json:"id"`
	Name_Feed_Back string            `json:"name_feed_back" validate:"required"`
	Room           entities.Room     `json:"room"`
	Description    string            `json:"description" validate:"required"`
	Category       entities.Category `json:"category"`
	User           entities.User     `json:"user"`
	Urls		   []string          `json:"url" validate:"required"`
}

package modelFeedback

import (
	"FM/src/entities"
	"time"
)

type GetFeedbackRes struct {
	ID             uint              `json:"id"`
	Name_Feed_Back string            `json:"name_feed_back" validate:"required"`
	Room           entities.Room     `json:"room"`
	Description    string            `json:"description" validate:"required"`
	Category       entities.Category `json:"category"`
	User           entities.User     `json:"user"`
	Urls           []string          `json:"url" validate:"required"`
}

type GetAllFeedbackRes struct {
	ID             uint      `json:"id"`
	Name_Feed_Back string    `json:"name_feed_back" validate:"required"`
	RoomID         uint      `json:"room_id" validate:"required,gte=1"`
	Description    string    `json:"description" validate:"required"`
	CategoryID     uint      `json:"category_id" validate:"required,gte=1"`
	UserID         uint      `json:"user_id" validate:"required,gte=1"`
	CategoryName   string    `json:"category_name" validate:"required"`
	RoomName       string    `json:"room_name" validate:"required"`
	Building       string    `json:"building_name" validate:"required"`
	Floor          string    `json:"floor_name" validate:"required"`
	Url            string    `json:"url" validate:"required"`
	CreateAt       time.Time `json:"created_at"`
	Urls		   []string  `json:"urls" validate:"required"`
}

type GetAllHistoryFeedbackRes struct {
	ID             uint              `json:"id"`
	Name_Feed_Back string            `json:"name_feed_back" validate:"required"`
	Room           entities.Room     `json:"room"`
	Description    string            `json:"description" validate:"required"`
	Category       entities.Category `json:"category"`
	User           entities.User     `json:"user"`
	Urls           []string          `json:"url" validate:"required"`
	CreatedAt      time.Time         `json:"created_at"`
	TimeStarted    time.Time         `json:"time_started"`
	TimeFinish     time.Time         `json:"time_finish"`
	Status         string            `json:"status"`
}

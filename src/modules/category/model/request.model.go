package model

type UpdateCategoryReq struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryReq struct {
	Name       string `json:"name"`
	FeedBackID int    `json:"feedback_id"`
}

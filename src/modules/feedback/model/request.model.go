package modelFeedback



type CreateFeedbackReq struct {
	NameFeedBack string `json:"name_feed_back"`
	Room         int64  `json:"room"`
}

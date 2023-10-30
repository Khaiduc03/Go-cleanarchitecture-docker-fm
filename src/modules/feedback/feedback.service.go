package feedback

import (
	modelFeedback "FM/src/modules/feedback/model"
	"context"
)

type FeedBackService interface {
	FindAll(ctx context.Context) ([]modelFeedback.GetAllFeedbackRes, error)
	FindById(ctx context.Context, id int) (modelFeedback.GetFeedbackRes, error)
	Create(ctx context.Context, model modelFeedback.CreateFeedbackReq) (string, error)
	History(ctx context.Context, user_id int) ([]modelFeedback.GetAllFeedbackRes, error)
	CheckCategory(ctx context.Context, category_id int) error
	CheckRoom(ctx context.Context, room_id int) error
	// Update(ctx context.Context, model modelFeedback.CreateFeedbackReq) (string, error)
	// Delete(ctx context.Context, id int) (string, error)
}

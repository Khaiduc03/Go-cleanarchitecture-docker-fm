package feedback

import (
	"FM/src/entities"
	modelFeedback "FM/src/modules/feedback/model"
	"context"
)

type FeedbackRepository interface {
	FindAll(ctx context.Context) ([]entities.FeedBack, error)
	FindById(ctx context.Context, id int) (entities.FeedBack, error)
	History(ctx context.Context, user_id int) ([]entities.FeedBack, error)
	Create(ctx context.Context, model modelFeedback.CreateFeedbackReq) (bool, error)
	CheckCategory(ctx context.Context, category_id int) (error)
	CheckRoom(ctx context.Context, room_id int) (error)
	// Update(ctx context.Context, model modelFeedback.CreateFeedbackReq) (string, error)
	// Delete(ctx context.Context, id int) (string, error)
}

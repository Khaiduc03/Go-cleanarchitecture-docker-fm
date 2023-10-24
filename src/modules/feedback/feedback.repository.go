package feedback

import (
	"FM/src/entities"
	modelFeedback "FM/src/modules/feedback/model"
	"context"
)

type FeedbackRepository interface {
	FindAll(ctx context.Context) ([]entities.FeedBack, error)
	FindById(ctx context.Context, id int) (entities.FeedBack, error)
	Create(ctx context.Context, model modelFeedback.CreateFeedbackReq) (string, error)
	Update(ctx context.Context, model modelFeedback.CreateFeedbackReq) (string, error)
	Delete(ctx context.Context, id int) (string, error)
}

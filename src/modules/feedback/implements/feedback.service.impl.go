package feedbackImpl

import (
	"FM/src/entities"
	"FM/src/modules/feedback"
	modelFeedback "FM/src/modules/feedback/model"
	"context"
)

type FeedbackServiceImpl struct {
	feedback.FeedbackRepository
}

func NewFeedbackServiceImpl(feedbackRepository *feedback.FeedbackRepository) feedback.FeedBackService {
	return &FeedbackServiceImpl{FeedbackRepository: *feedbackRepository}
}

func (feedbackService *FeedbackServiceImpl) FindAll(ctx context.Context) ([]entities.FeedBack, error) {
	return feedbackService.FeedbackRepository.FindAll(ctx)
}

func (feedbackService *FeedbackServiceImpl) FindById(ctx context.Context, id int) (entities.FeedBack, error) {
	return feedbackService.FeedbackRepository.FindById(ctx, id)
}

func (feedbackService *FeedbackServiceImpl) Create(ctx context.Context, req modelFeedback.CreateFeedbackReq) (string, error) {
	res, err := feedbackService.FeedbackRepository.Create(ctx, req)
	if err != nil && !res {
		return "Create feedback failed", err
	}
	return "Create feedback success", nil
}

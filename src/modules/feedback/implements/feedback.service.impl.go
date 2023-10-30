package feedbackImpl

import (
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

func (feedbackService *FeedbackServiceImpl) FindAll(ctx context.Context) ([]modelFeedback.GetAllFeedbackRes, error) {
	return feedbackService.FeedbackRepository.FindAll(ctx)
}

func (feedbackService *FeedbackServiceImpl) FindById(ctx context.Context, id int) (modelFeedback.GetFeedbackRes, error) {
	return feedbackService.FeedbackRepository.FindById(ctx, id)
}

func (feedbackService *FeedbackServiceImpl) Create(ctx context.Context, req modelFeedback.CreateFeedbackReq) (string, error) {

	res, err := feedbackService.FeedbackRepository.Create(ctx, req)
	if err != nil && !res {
		return "Create feedback failed", err
	}
	return "Create feedback success", nil
}

func (feedbackService *FeedbackServiceImpl) History(ctx context.Context, user_id int) ([]modelFeedback.GetAllFeedbackRes, error) {
	return feedbackService.FeedbackRepository.History(ctx, user_id)
}

func (feedbackService *FeedbackServiceImpl) CheckCategory(ctx context.Context, category_id int) error {
	return feedbackService.FeedbackRepository.CheckCategory(ctx, category_id)
}

func (feedbackService *FeedbackServiceImpl) CheckRoom(ctx context.Context, room_id int) error {
	return feedbackService.FeedbackRepository.CheckRoom(ctx, room_id)
}

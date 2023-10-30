package feedbackImpl

import (
	"FM/src/entities"
	"FM/src/modules/feedback"
	modelFeedback "FM/src/modules/feedback/model"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type FeedbackRepositoryImpl struct {
	*gorm.DB
}

func NewFeedbackRepositoryImpl(DB *gorm.DB) feedback.FeedbackRepository {
	return &FeedbackRepositoryImpl{DB: DB}
}

// FindAll implements feedback.FeedbackRepository.
func (feedbackRepository *FeedbackRepositoryImpl) FindAll(ctx context.Context) ([]entities.FeedBack, error) {

	var feedbacks []entities.FeedBack
	err := feedbackRepository.DB.Find(&feedbacks).Error

	return feedbacks, err
}

// FindById implements feedback.FeedbackRepository.
func (feedbackRepository *FeedbackRepositoryImpl) FindById(ctx context.Context, id int) (entities.FeedBack, error) {

	var feedback entities.FeedBack
	isExist := feedbackRepository.DB.Where("id = ?", id).Find(&feedback)
	if isExist.RowsAffected == 0 {
		return feedback, errors.New("feedback not found")
	}
	return feedback, nil
}

func (feedbackRepository *FeedbackRepositoryImpl) Create(ctx context.Context, req modelFeedback.CreateFeedbackReq) (bool, error) {
	var feedback entities.FeedBack

	var image entities.Image

	feedback = entities.FeedBack{UserID: uint(req.UserID), NameFeedBack: req.Name_Feed_Back, Description: req.Description, CategoryID: uint(req.CategoryID), RoomID: uint(req.RoomID)}

	err := feedbackRepository.DB.Create(&feedback)
	//err := feedbackRepository.DB.Create(&req).Error
	if err.Error != nil {
		return false, err.Error
	}
	for _, url := range req.Urls {
		image = entities.Image{FeedbackID: feedback.ID, Url: url}
		errImg := feedbackRepository.DB.Create(&image).Error
		if errImg != nil {
			return false, errImg
		}
		fmt.Println("image", image)
	}

	return true, nil
}

func (feedbackRepository *FeedbackRepositoryImpl) History(ctx context.Context, user_id int) ([]entities.FeedBack, error) {
	var feedbacks []entities.FeedBack
	err := feedbackRepository.DB.Where("user_id = ?", user_id).Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (feedbackRepository *FeedbackRepositoryImpl) CheckCategory(ctx context.Context, category_id int) ( error) {
	var category entities.Category
	isExistCategory := feedbackRepository.DB.Where("id = ?", category_id).Find(&category)
	if isExistCategory.RowsAffected == 0 {
		return  errors.New("category not found")
	}
	return  nil
}

func (feedbackRepository *FeedbackRepositoryImpl) CheckRoom(ctx context.Context, room_id int) ( error) {
	var room entities.Room
	isExistRoom := feedbackRepository.DB.Where("id = ?", room_id).Find(&room)
	if isExistRoom.RowsAffected == 0 {
		return  errors.New("room not found")
	}
	return nil
}

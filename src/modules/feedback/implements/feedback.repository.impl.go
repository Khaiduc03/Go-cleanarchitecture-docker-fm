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
func (feedbackRepository *FeedbackRepositoryImpl) FindAll(ctx context.Context) ([]modelFeedback.GetAllFeedbackRes, error) {

	//var feedbacks []modelFeedback.GetAllFeedbackRes
	var result []modelFeedback.GetAllFeedbackRes
	err := feedbackRepository.DB.Raw(`SELECT f.*, 
	c.category_name,
	r.room_name,
	r.floor,
	r.building,
	u.name,
	u.url
	FROM "FEEDBACK" f 
	LEFT JOIN "CATEGORY" c ON f.category_id = c.id 
	LEFT JOIN "ROOM" r ON f.room_id = r.id
	LEFT JOIN "USER" u ON f.user_id = u.id
	WHERE f.reciever_id IS NULL`).Scan(&result).Error

	if err != nil {
		return nil, err
	}
	for i, feedback := range result {
		var images []entities.Image
		var urls []string
		err = feedbackRepository.DB.Select("url").Where("feedback_id = ?", feedback.ID).Find(&images).Error

		if err != nil {
			return nil, err
		}
		for _, image := range images {

			urls = append(urls, image.Url)
		}
		result[i].Urls = urls
	}

	return result, nil
}

// FindById implements feedback.FeedbackRepository.
func (feedbackRepository *FeedbackRepositoryImpl) FindById(ctx context.Context, id int) (modelFeedback.GetFeedbackRes, error) {

	var feedback entities.FeedBack
	err := feedbackRepository.DB.Where("id = ?", id).First(&feedback).Error
	if err != nil {
		return modelFeedback.GetFeedbackRes{}, err
	}

	var images []entities.Image
	var urls []string
	err = feedbackRepository.DB.Select("url").Where("feedback_id = ?", id).Find(&images).Error
	if err != nil {
		return modelFeedback.GetFeedbackRes{}, err
	}
	for _, image := range images {

		urls = append(urls, image.Url)
	}

	var user entities.User
	err = feedbackRepository.DB.Where("id = ?", feedback.UserID).Find(&user).Error
	if err != nil {
		return modelFeedback.GetFeedbackRes{}, err
	}

	var room entities.Room
	err = feedbackRepository.DB.Where("id = ?", feedback.RoomID).Find(&room).Error
	if err != nil {
		return modelFeedback.GetFeedbackRes{}, err
	}
	var category entities.Category
	err = feedbackRepository.DB.Where("id = ?", feedback.CategoryID).Find(&category).Error
	if err != nil {
		return modelFeedback.GetFeedbackRes{}, err
	}
	res := modelFeedback.GetFeedbackRes{
		ID:             feedback.ID,
		Name_Feed_Back: feedback.NameFeedBack,
		Room:           room,
		Description:    feedback.Description,
		Category:       category,
		User:           user,
		Urls:           urls,
	}

	return res, nil
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

func (feedbackRepository *FeedbackRepositoryImpl) History(ctx context.Context, user_id int) ([]modelFeedback.GetAllHistoryFeedbackRes, error) {
	var feedbacks []entities.FeedBack
	err := feedbackRepository.DB.Where("user_id = ?", user_id).Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}
	var user entities.User
	err = feedbackRepository.DB.Where("id = ?", user_id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	//image
	var images []entities.Image
	var urls []string

	var result []modelFeedback.GetAllHistoryFeedbackRes
	for _, feedback := range feedbacks {
		//room
		var room entities.Room
		err = feedbackRepository.DB.Where("id = ?", feedback.RoomID).Find(&room).Error
		if err != nil {
			return nil, err
		}
		//category
		var category entities.Category
		err = feedbackRepository.DB.Where("id = ?", feedback.CategoryID).Find(&category).Error
		if err != nil {
			return nil, err
		}

		err = feedbackRepository.DB.Select("url").Where("feedback_id = ?", feedback.ID).Find(&images).Error
		if err != nil {
			return nil, err
		}
		for _, image := range images {

			urls = append(urls, image.Url)
		}

		f := modelFeedback.GetAllHistoryFeedbackRes{
			ID:             feedback.ID,
			Name_Feed_Back: feedback.NameFeedBack,
			Description:    feedback.Description,
			Category:       category,
			Room:           room,
			User:           user,
			Urls:           urls,
			CreatedAt:      feedback.CreatedAt,
			TimeStarted:    feedback.TimeStarted,
			TimeFinish:     feedback.TimeFinish,
			Status:         feedback.Status,
		}

		result = append(result, f)
	}
	return result, nil
}

func (feedbackRepository *FeedbackRepositoryImpl) CheckCategory(ctx context.Context, category_id int) error {
	var category entities.Category
	isExistCategory := feedbackRepository.DB.Where("id = ?", category_id).Find(&category)
	if isExistCategory.RowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (feedbackRepository *FeedbackRepositoryImpl) CheckRoom(ctx context.Context, room_id int) error {
	var room entities.Room
	isExistRoom := feedbackRepository.DB.Where("id = ?", room_id).Find(&room)
	if isExistRoom.RowsAffected == 0 {
		return errors.New("room not found")
	}
	return nil
}

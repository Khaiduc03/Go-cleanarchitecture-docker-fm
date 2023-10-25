package roomImpl

import (
	"FM/src/entities"
	room "FM/src/modules/Room"
	modelRoom "FM/src/modules/Room/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type RoomRepositoryImpl struct {
	*gorm.DB
}

func NewRoomRepositoryImpl(DB *gorm.DB) room.RoomRepository {
	return &RoomRepositoryImpl{DB: DB}
}

func (roomRepository *RoomRepositoryImpl) FindAll(ctx context.Context) ([]entities.Room, error) {
	var rooms []entities.Room
	err := roomRepository.DB.Find(&rooms).Error

	return rooms, err
}

func (roomRepository *RoomRepositoryImpl) FindById(ctx context.Context, id int) (entities.Room, error) {
	var room entities.Room
	err := roomRepository.DB.Where("id = ?", id).Find(&room).Error

	return room, err
}

func (roomRepository *RoomRepositoryImpl) Create(ctx context.Context, req modelRoom.CreateRoomReq) (bool, error) {
	var room entities.Room
	isExist := roomRepository.DB.Where("room_name = ? ", req.Room_Name).Find(&room)
	if isExist.RowsAffected != 0 {
		return false, nil
	}
	fmt.Println(req)
	room = entities.Room{RoomName: req.Room_Name, Floor: req.Floor, Building: req.Building, Status: 1}
	err := roomRepository.DB.Create(&room).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func (roomRepository *RoomRepositoryImpl) Update(ctx context.Context, req modelRoom.UpdateRoomReq) (bool, error) {
	var room entities.Room
	result := roomRepository.DB.Where("id = ?", req.ID).First(&room)
	if result.RowsAffected == 0 {
		return false, nil
	}
	room.RoomName = req.Room_Name
	room.Floor = req.Floor
	room.Building = req.Building
	room.Status = req.Status

	err := roomRepository.DB.Save(&room).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (roomRepository *RoomRepositoryImpl) Delete(ctx context.Context, id int) (bool, error) {
	var room entities.Room
	result := roomRepository.DB.Where("id = ?", id).First(&room)
	if result.Error != nil {
		return false, result.Error
	}
	
	if result.RowsAffected == 0 {
		return false, nil
	}
	
	err := roomRepository.DB.Delete(&room).Error
	if err != nil {
		return false, err
	}
	
	return true, nil
}

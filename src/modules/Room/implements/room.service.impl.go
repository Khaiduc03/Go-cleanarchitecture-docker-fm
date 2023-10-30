package roomImpl

import (
	"FM/src/entities"
	room "FM/src/modules/Room"
	modelRoom "FM/src/modules/Room/model"
	"context"
)

type RoomServiceImpl struct {
	room.RoomRepository
}

func NewRoomServiceImpl(roomRepository *room.RoomRepository) room.RoomService {
	return &RoomServiceImpl{RoomRepository: *roomRepository}
}

func (roomService *RoomServiceImpl) FindAll(ctx context.Context, room_name string) ([]entities.Room, error) {
	if room_name != "" {
		return roomService.RoomRepository.FindByName(ctx, room_name)
	}
	return roomService.RoomRepository.FindAll(ctx)
}

func (roomService *RoomServiceImpl) FindById(ctx context.Context, id int) (entities.Room, error) {
	return roomService.RoomRepository.FindById(ctx, id)
}

func (roomService *RoomServiceImpl) Create(ctx context.Context, req modelRoom.CreateRoomReq) (string, error) {

	response, err := roomService.RoomRepository.Create(ctx, req)
	if err != nil && !response {
		return "Create room failed", err
	}
	return "Create room success", nil
}

func (roomService *RoomServiceImpl) Update(ctx context.Context, req modelRoom.UpdateRoomReq) (string, error) {

	response, err := roomService.RoomRepository.Update(ctx, req)
	if err != nil && !response {
		return "Update room failed", err
	}
	return "Update room success", nil
}

func (roomService *RoomServiceImpl) Delete(ctx context.Context, id int) (string, error) {

	response, err := roomService.RoomRepository.Delete(ctx, id)
	if err != nil && !response {
		return "Delete room failed", err
	}
	return "Delete room success", nil
}

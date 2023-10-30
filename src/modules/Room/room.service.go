package room

import (
	"FM/src/entities"
	modelRoom "FM/src/modules/Room/model"
	"context"
)


type RoomService interface {
	FindAll(ctx context.Context, room_name string ) ([]entities.Room, error)
	FindById(ctx context.Context, id int) (entities.Room, error)

	Create(ctx context.Context, req modelRoom.CreateRoomReq) (string, error)
	Update(ctx context.Context, req modelRoom.UpdateRoomReq) (string, error)
	Delete(ctx context.Context, id int) (string, error)
}

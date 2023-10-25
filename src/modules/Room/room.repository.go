package room

import (
	"FM/src/entities"
	modelRoom "FM/src/modules/Room/model"
	"context"
)

type RoomRepository interface {
	FindAll(ctx context.Context) ([]entities.Room, error)
	FindById(ctx context.Context, id int) (entities.Room, error)
	Create(ctx context.Context, req modelRoom.CreateRoomReq) (bool, error)
	Update(ctx context.Context, req modelRoom.UpdateRoomReq) (bool, error)
	Delete(ctx context.Context, id int) (bool, error)
}



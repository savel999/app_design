package storage

import (
	"context"
	"errors"
	"fmt"

	domainerrors "github.com/savel999/app_design/internal/domain/errors"
	"github.com/savel999/app_design/internal/domain/models"
	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type RoomsRepository struct {
	roomsStore storage.Rooms
}

func NewRoomsRepository(roomsStore storage.Rooms) *RoomsRepository {
	return &RoomsRepository{roomsStore: roomsStore}
}

func (r *RoomsRepository) GetByID(ctx context.Context, id int) (models.Room, error) {
	room, err := r.roomsStore.GetByID(ctx, id)
	switch {
	case errors.Is(err, storage.ErrNotFound):
		return models.Room{}, domainerrors.ErrNotFound
	case err != nil:
		return models.Room{}, fmt.Errorf("can't get room by id: %w", err)
	default:
		return mapRoomToModel(room), nil
	}
}

func (r *RoomsRepository) Create(ctx context.Context, q models.Room) (models.Room, error) {
	in := dto.RoomRaw{
		Name:    q.Name,
		HotelID: q.HotelID,
		Count:   q.Count,
		Price:   q.Price,
	}

	room, err := r.roomsStore.Create(ctx, in)
	if err != nil {
		return models.Room{}, fmt.Errorf("can't create room: %w", err)
	}

	return mapRoomToModel(room), nil
}

func mapRoomToModel(u dto.RoomRaw) models.Room {
	return models.Room{ID: u.ID, HotelID: u.HotelID, Name: u.Name, Count: u.Count, Price: u.Price}
}

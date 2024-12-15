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

type HotelsRepository struct {
	hotelsStore storage.Hotels
}

func NewHotelsRepository(hotelsStore storage.Hotels) *HotelsRepository {
	return &HotelsRepository{hotelsStore: hotelsStore}
}

func (r *HotelsRepository) Create(ctx context.Context, q models.Hotel) (models.Hotel, error) {
	hotel, err := r.hotelsStore.Create(ctx, dto.HotelRaw{Name: q.Name, CheckIn: q.CheckIn, CheckOut: q.CheckOut})
	if err != nil {
		return models.Hotel{}, fmt.Errorf("can't create hotel: %w", err)
	}

	return mapHotelToModel(hotel), nil
}

func (r *HotelsRepository) GetByID(ctx context.Context, id int) (models.Hotel, error) {
	room, err := r.hotelsStore.GetByID(ctx, id)
	switch {
	case errors.Is(err, storage.ErrNotFound):
		return models.Hotel{}, domainerrors.ErrNotFound
	case err != nil:
		return models.Hotel{}, fmt.Errorf("can't get hotel by id: %w", err)
	default:
		return mapHotelToModel(room), nil
	}
}

func mapHotelToModel(h dto.HotelRaw) models.Hotel {
	return models.Hotel{ID: h.ID, Name: h.Name}
}

package memory

import (
	"context"
	"sync"

	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type rooms struct {
	mu     *sync.RWMutex
	data   map[int]dto.RoomRaw
	lastID int
}

func NewRoomsStorage() storage.Rooms {
	return &rooms{mu: &sync.RWMutex{}, data: make(map[int]dto.RoomRaw)}
}

func (s *rooms) GetByID(_ context.Context, id int) (dto.RoomRaw, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, v := range s.data {
		if v.ID == id {
			return v, nil
		}
	}

	return dto.RoomRaw{}, storage.ErrNotFound
}

func (s *rooms) Create(_ context.Context, room dto.RoomRaw) (dto.RoomRaw, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++

	room.ID = s.lastID
	s.data[s.lastID] = room

	return room, nil
}

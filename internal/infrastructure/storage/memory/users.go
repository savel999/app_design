package memory

import (
	"context"
	"sync"

	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type users struct {
	mu     *sync.RWMutex
	data   map[int]dto.UserRaw
	lastID int
}

func NewUsersStorage() storage.Users {
	return &users{mu: &sync.RWMutex{}, data: make(map[int]dto.UserRaw)}
}

func (s *users) GetByEmail(_ context.Context, email string) (dto.UserRaw, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, v := range s.data {
		if v.Email == email {
			return v, nil
		}
	}

	return dto.UserRaw{}, storage.ErrNotFound
}

func (s *users) Create(_ context.Context, user dto.UserRaw) (dto.UserRaw, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++

	user.ID = s.lastID
	s.data[s.lastID] = user

	return user, nil
}

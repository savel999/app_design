package memory

import (
	"context"
	"sync"

	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type payments struct {
	mu     *sync.RWMutex
	data   map[int]dto.PaymentRaw
	lastID int
}

func NewPaymentsStorage() storage.Payments {
	return &payments{mu: &sync.RWMutex{}, data: make(map[int]dto.PaymentRaw)}
}

func (s *payments) Create(_ context.Context, user dto.PaymentRaw) (dto.PaymentRaw, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++

	user.ID = s.lastID
	s.data[s.lastID] = user

	return user, nil
}

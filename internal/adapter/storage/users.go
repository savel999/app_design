package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/savel999/app_design/internal/domain/models"
	"github.com/savel999/app_design/internal/domain/repos"
	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type UsersRepository struct {
	usersStore storage.Users
}

func NewUsersRepository(store storage.Users) *UsersRepository {
	return &UsersRepository{usersStore: store}
}

func (r *UsersRepository) GetOrCreate(ctx context.Context, in repos.GetOrCreateQuery) (models.User, error) {
	user, err := r.usersStore.GetByEmail(ctx, in.Email)
	switch {
	case errors.Is(err, storage.ErrNotFound):
		newUser, createErr := r.Create(ctx, models.User{Email: in.Email})
		if createErr != nil {
			return models.User{}, fmt.Errorf("can't create user: %w", createErr)
		}

		return newUser, nil
	case err != nil:
		return models.User{}, fmt.Errorf("can't get user by email: %w", err)
	default:
		return mapUserToModel(user), nil
	}
}

func (r *UsersRepository) Create(ctx context.Context, q models.User) (models.User, error) {
	user, err := r.usersStore.Create(ctx, dto.UserRaw{Email: q.Email})
	if err != nil {
		return models.User{}, fmt.Errorf("can't create room: %w", err)
	}

	return mapUserToModel(user), nil
}

func mapUserToModel(u dto.UserRaw) models.User {
	return models.User{ID: u.ID, Email: u.Email}
}

package web

import (
	"github.com/savel999/app_design/internal/adapter/storage"
	"github.com/savel999/app_design/internal/app/web/registry"
	"github.com/savel999/app_design/internal/app/web/usecase"
	"github.com/savel999/app_design/internal/domain/services"
	"github.com/savel999/app_design/internal/infrastructure/logger"
	"github.com/savel999/app_design/internal/infrastructure/storage/memory"
)

func NewContainer(
	config *Config,
) *registry.Container {
	log := logger.New(config.LogLevel)
	repositories := newRepositories()

	return &registry.Container{
		Logger:       log,
		Usecases:     newUsecases(repositories, newServices(repositories)),
		Repositories: repositories,
	}
}

func newUsecases(repositories *registry.Repositories, services *registry.Services) *registry.Usecases {
	return &registry.Usecases{
		CreateOrder: usecase.NewCreateOrder(
			repositories.Users,
			repositories.Hotels,
			repositories.Rooms,
			repositories.Bookings,
			services.Orders,
		),
	}
}

func newRepositories() *registry.Repositories {
	return &registry.Repositories{
		Users:    storage.NewUsersRepository(memory.NewUsersStorage()),
		Rooms:    storage.NewRoomsRepository(memory.NewRoomsStorage()),
		Hotels:   storage.NewHotelsRepository(memory.NewHotelsStorage()),
		Bookings: storage.NewBookingsRepository(memory.NewBookingsStorage()),
		Orders:   storage.NewOrdersRepository(memory.NewOrdersStorage()),
		Payments: storage.NewPaymentsRepository(memory.NewPaymentsStorage()),
	}
}

func newServices(r *registry.Repositories) *registry.Services {
	return &registry.Services{
		Orders: services.NewOrdersService(
			r.Bookings,
			r.Orders,
			r.Payments,
		),
	}
}

package registry

import (
	"github.com/savel999/app_design/internal/app/web/usecase"
	"github.com/savel999/app_design/internal/domain/repos"
	"github.com/savel999/app_design/internal/domain/services"
	"github.com/savel999/app_design/internal/infrastructure/logger"
)

type Container struct {
	Logger       logger.Logger
	Usecases     *Usecases
	Repositories *Repositories
	Services     *Services
}

type Usecases struct {
	CreateOrder usecase.CreateOrder
}

type Repositories struct {
	Users    repos.UsersRepo
	Hotels   repos.HotelsRepo
	Rooms    repos.RoomsRepo
	Bookings repos.BookingsRepo
	Orders   repos.OrdersRepo
	Payments repos.PaymentsRepo
}

type Services struct {
	Orders services.OrdersService
}

func (cnt *Container) Clean() {
	//clean resources
}

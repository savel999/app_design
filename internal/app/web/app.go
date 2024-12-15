package web

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/savel999/app_design/internal/app/web/registry"
	"github.com/savel999/app_design/internal/app/web/server"
	"github.com/savel999/app_design/internal/domain/models"
)

func Run(config *Config) error {
	cnt := NewContainer(config)

	cnt.Logger.Debug("init config", slog.Any("config", config))

	s := InitServer(cnt)

	if config.WithDemoData {
		InitDemoData(cnt)
	}

	beforeShutdown := func() {
		cnt.Clean()
	}

	if err := s.ListenAndServe(config.ServerAddr, beforeShutdown); err != nil {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	return nil
}

func InitServer(cnt *registry.Container) *server.Server {
	s := server.NewServer(cnt.Logger)

	s.InitMiddlewares()
	s.InitProbes()
	s.InitRoutes(cnt)

	return s
}

func InitDemoData(cnt *registry.Container) {
	users := []models.User{{Email: "test@test.com"}, {Email: "demo@demo.com"}}

	for _, user := range users {
		cnt.Repositories.Users.Create(context.Background(), user)
	}

	parseTime := func(onlyTime string) time.Time {
		v, _ := time.Parse(time.TimeOnly, onlyTime)

		return v
	}

	hotels := []models.Hotel{
		{Name: "Marina Resort", CheckIn: parseTime("14:00:00"), CheckOut: parseTime("12:00:00")},
		{Name: "Ozkaymak Hotel", CheckIn: parseTime("13:00:00"), CheckOut: parseTime("11:00:00")},
	}

	for _, hotelItem := range hotels {
		hotel, err := cnt.Repositories.Hotels.Create(context.Background(), hotelItem)
		if err == nil {
			rooms := []models.Room{
				{Name: "lux", Count: 1, Price: 15000},
				{Name: "standard", Count: 2, Price: 5000},
			}

			for _, roomItem := range rooms {
				roomItem.HotelID = hotel.ID

				cnt.Repositories.Rooms.Create(context.Background(), roomItem)
			}
		}
	}
}

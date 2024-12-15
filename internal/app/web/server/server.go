package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/savel999/app_design/internal/app/web/registry"
	"github.com/savel999/app_design/internal/infrastructure/logger"
	"github.com/savel999/app_design/internal/presentation/rest/handlers"
)

const graphqlHTTPPath = "/graphql"

type Server struct {
	router *chi.Mux
	logger logger.Logger
}

func NewServer(logger logger.Logger) *Server {
	return &Server{
		router: chi.NewRouter(),
		logger: logger,
	}
}

func (s *Server) GetRouter() *chi.Mux {
	return s.router
}

func (s *Server) InitMiddlewares() *Server {
	const (
		maxAgePreflightRequest = 300
		allowedContentType     = "application/json"
	)

	s.router.Use(middleware.AllowContentType(allowedContentType))
	s.router.Use(middleware.SetHeader("Content-Type", allowedContentType))
	s.router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           maxAgePreflightRequest,
	}))

	return s
}

func (s *Server) InitProbes() {
	const readinessPath = "/readiness"

	s.router.Get(readinessPath, func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write([]byte("i'm ready")); err != nil {
			s.logger.Error(fmt.Sprintf("failed to handle %s location", readinessPath), logger.ErrorAttr(err))
		}
	})
}

func (s *Server) InitRoutes(cnt *registry.Container) {
	h := handlers.New(cnt.Logger, cnt.Usecases)

	s.router.Post("/orders", h.CreateOrder())
	s.router.HandleFunc("/*", h.ErrorHandler(http.StatusMethodNotAllowed, "route not found"))
}

func (s *Server) ListenAndServe(addr string, shutdownFn func()) error {
	srv := http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	srv.RegisterOnShutdown(shutdownFn)

	exitCh := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			s.logger.Error("failed to shutdown server", logger.ErrorAttr(err))
		}

		close(exitCh)
	}()

	s.logger.Info("ListenAndServe", slog.String("addr", addr))

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to ListenAndServe: %w", err)
	}

	<-exitCh

	return nil
}

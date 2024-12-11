package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/savel999/app_design/internal/infrastructure/logger"
)

type Handler struct {
	logger logger.Logger
}

type responseError struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"error,omitempty"`
}

func New(logger logger.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) Handle(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				h.logger.ErrorContext(r.Context(), "handler panic", slog.Any("err", rvr))

				h.ErrorHandler(http.StatusInternalServerError, "unexpectable error")(w, r)
			}
		}()

		fn(w, r)
	}
}

func (h *Handler) ErrorHandler(
	code int, message string, errors ...error,
) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)

		resp := responseError{Code: http.StatusText(code), Message: message}

		if len(errors) > 0 {
			resp.Errors = make([]string, 0, len(errors))

			for _, err := range errors {
				resp.Errors = append(resp.Errors, err.Error())
			}
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			h.logger.ErrorContext(r.Context(), "failed to encode response", logger.ErrorAttr(err))
		}
	}
}

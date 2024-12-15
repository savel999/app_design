package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/savel999/app_design/internal/app/web/registry"
	"github.com/savel999/app_design/internal/infrastructure/logger"
)

type Handler struct {
	logger   logger.Logger
	usecases *registry.Usecases
}

type responseError struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"error,omitempty"`
}

func New(logger logger.Logger, usecases *registry.Usecases) *Handler {
	return &Handler{logger: logger, usecases: usecases}
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

func (h *Handler) decodeRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, in any) bool {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&in); err != nil {
		h.logger.WarnContext(ctx, "failed to decode request body", logger.ErrorAttr(err))
		h.ErrorHandler(http.StatusBadRequest, "can't parse request body")(w, r)

		return false
	}

	return true
}

func (h *Handler) encodeResponse(ctx context.Context, w http.ResponseWriter, r *http.Request, status int, out any) {
	if err := json.NewEncoder(w).Encode(out); err != nil {
		h.logger.ErrorContext(ctx, "failed to encode response", logger.ErrorAttr(err))
		h.ErrorHandler(http.StatusInternalServerError, "can't parse request body")(w, r)

		return
	}

	w.WriteHeader(status)
}

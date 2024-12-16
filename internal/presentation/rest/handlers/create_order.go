package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/savel999/app_design/internal/infrastructure/logger"
	"github.com/savel999/app_design/internal/presentation/rest/dto"
)

func (h *Handler) CreateOrder() http.HandlerFunc {
	return h.Handle(func(w http.ResponseWriter, r *http.Request) {
		ctx, in := r.Context(), dto.CreateOrderRequest{}
		if !h.decodeRequest(ctx, w, r, &in) {
			return
		}

		out, err := h.usecases.CreateOrder(ctx, in)
		if err != nil {
			var validationErrors *dto.ValidationErrors

			switch {
			case errors.As(err, &validationErrors):
				h.logger.WarnContext(
					ctx,
					"failed to create order(validation errors)",
					logger.ErrorAttr(err), slog.Any("in", in),
				)

				h.ErrorHandler(http.StatusUnprocessableEntity, validationErrors.Message, validationErrors.Errors...)(w, r)
			default:
				h.logger.ErrorContext(
					ctx, "failed to create order", logger.ErrorAttr(err), slog.Any("in", in),
				)

				h.ErrorHandler(http.StatusInternalServerError, "can't create order")(w, r)
			}

			return
		}

		w.WriteHeader(http.StatusCreated)

		h.encodeResponse(ctx, w, r, out)
	})
}

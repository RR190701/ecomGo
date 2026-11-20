package orders

import (
	"log/slog"
	"net/http"

	"github.com/RR190701/ecomGo/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {

	return &handler{
		service: s,
	}
}

func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var tempOrder createOrderParams
	err := json.Read(r, &tempOrder)

	if err != nil {
		slog.Error("Request Deocde error", "error", err.Error())
		http.Error(w, "Request Decode error", http.StatusBadRequest)
		return
	}

	order, err := h.service.CreateOrder(r.Context(), tempOrder)
	if err != nil {
		slog.Error("Failed to create order", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, order)

}

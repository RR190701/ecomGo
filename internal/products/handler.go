package products

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

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error("Failed to list products", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products := []string{"product1", "product2", "product3"}
	json.Write(w, http.StatusOK, products)

}

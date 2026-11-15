package products

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/RR190701/ecomGo/internal/json"
	"github.com/go-chi/chi/v5"
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

	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error("Failed to list products", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, products)

}

func (h *handler) ListOneProduct(w http.ResponseWriter, r *http.Request) {

	param := chi.URLParam(r, "id")

	if param == "" {
		slog.Error("Missing product ID in request")
		http.Error(w, "Missing product ID in request", http.StatusBadRequest)
		return
	}
	productId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		slog.Error("Invalid product ID in request", "error", err.Error())
		http.Error(w, "Invalid product ID in request", http.StatusBadRequest)
		return
	}

	products, err := h.service.ListOneProduct(r.Context(), productId)
	if err != nil {
		slog.Error("Failed to list products", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, products)

}

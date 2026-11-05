package products

import (
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
	products := []string{"product1", "product2", "product3"}
	json.Write(w, http.StatusOK, products)

}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ericoalmeida/go-orders_clean-arch/internal/usecases"
)

type GetAllOrdersHandler struct {
	usecase *usecases.GetAllOrdersUsecase
}

func NewGetAllOrdersHandler(input *usecases.GetAllOrdersUsecase) *GetAllOrdersHandler {
	return &GetAllOrdersHandler{usecase: input}
}

func (handler *GetAllOrdersHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := handler.usecase.ListAll()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

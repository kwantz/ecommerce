package controller

import (
	"net/http"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
	"github.com/kwantz/ecommerce/services/order/internal/utility"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (controller *PingController) PingHandler(w http.ResponseWriter, r *http.Request) {
	response := entity.PingResponse{
		Message: "Hello World from Order Service",
	}

	utility.ResponseJSON(w, response)
}

package controller

import (
	"net/http"

	"github.com/kwantz/ecommerce/services/account/internal/entity"
	"github.com/kwantz/ecommerce/services/account/internal/utility"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (controller *PingController) PingHandler(w http.ResponseWriter, r *http.Request) {
	response := entity.PingResponse{
		Message: "Hello World from Account Service",
	}

	utility.ResponseJSON(w, response)
}

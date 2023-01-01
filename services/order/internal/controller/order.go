package controller

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/kwantz/ecommerce/services/order/internal/utility"
)

type OrderController struct {
	orderUsecase   OrderUsecase
	accountUsecase AccountUsecase
}

func NewOrderController(orderUsecase OrderUsecase, accountUsecase AccountUsecase) *OrderController {
	return &OrderController{
		orderUsecase:   orderUsecase,
		accountUsecase: accountUsecase,
	}
}

func (controller *OrderController) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	operation := "OrderController.CreateOrderHandler"

	auth := r.Header.Get("Authorization")
	if !strings.Contains(auth, "Bearer") {
		log.Printf("[%s] invalid authorization header", operation)
		utility.ResponseErrorJSON(w, errors.New("bad request"))
		return
	}

	accountID, err := controller.accountUsecase.Authorization(r.Context(), auth)
	if err != nil {
		log.Printf("[%s] failed authorize, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	response, err := controller.orderUsecase.OrderProduct(r.Context(), auth, accountID)
	if err != nil {
		log.Printf("[%s] failed create order from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

package controller

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
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

func (controller *OrderController) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	operation := "OrderController.GetOrdersHandler"

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

	response, err := controller.orderUsecase.GetOrders(r.Context(), accountID)
	if err != nil {
		log.Printf("[%s] failed get order from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

func (controller *OrderController) GetOrderDetailHandler(w http.ResponseWriter, r *http.Request) {
	operation := "OrderController.GetOrderDetailHandler"

	invoice := chi.URLParam(r, "invoice")

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

	response, err := controller.orderUsecase.GetOrderDetail(r.Context(), accountID, invoice)
	if err != nil {
		log.Printf("[%s] failed get order detail from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

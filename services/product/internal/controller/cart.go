package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/kwantz/ecommerce/services/product/internal/entity"
	"github.com/kwantz/ecommerce/services/product/internal/utility"
)

type CartController struct {
	cartUsecase    CartUsecase
	accountUsecase AccountUsecase
}

func NewCartController(cartUsecase CartUsecase, accountUsecase AccountUsecase) *CartController {
	return &CartController{
		cartUsecase:    cartUsecase,
		accountUsecase: accountUsecase,
	}
}

func (controller *CartController) AddProductToCartHandler(w http.ResponseWriter, r *http.Request) {
	operation := "CartController.AddProductToCartHandler"

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

	request := entity.CartRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("[%s] failed decode request body, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	request.AccountID = accountID
	response, err := controller.cartUsecase.AddProductToCart(r.Context(), request)
	if err != nil {
		log.Printf("[%s] failed add item to cart from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

func (controller *CartController) RemoveProductFromCartHandler(w http.ResponseWriter, r *http.Request) {
	operation := "CartController.RemoveProductFromCartHandler"

	auth := r.Header.Get("Authorization")
	if !strings.Contains(auth, "Bearer") {
		log.Printf("[%s] invalid authorization header", operation)
		utility.ResponseErrorJSON(w, errors.New("bad request"))
		return
	}

	if _, err := controller.accountUsecase.Authorization(r.Context(), auth); err != nil {
		log.Printf("[%s] failed authorize, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	strCartID := chi.URLParam(r, "cartID")
	cartID, err := strconv.ParseInt(strCartID, 10, 64)
	if err != nil {
		log.Printf("[%s] failed parse cart ID to int, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	response, err := controller.cartUsecase.RemoveProductFromCart(r.Context(), cartID)
	if err != nil {
		log.Printf("[%s] failed add item to cart from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

func (controller *CartController) GetCartProductsOrderHandler(w http.ResponseWriter, r *http.Request) {
	operation := "CartController.GetCartProductsOrderHandler"

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

	response, err := controller.cartUsecase.GetCartProductsByAccountID(r.Context(), accountID)
	if err != nil {
		log.Printf("[%s] failed get cart products by account ID from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

func (controller *CartController) DeleteCartProductsOrderHandler(w http.ResponseWriter, r *http.Request) {
	operation := "CartController.DeleteCartProductsOrderHandler"

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

	response, err := controller.cartUsecase.DeleteCartProductsByAccountID(r.Context(), accountID)
	if err != nil {
		log.Printf("[%s] failed delete cart products by account ID from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

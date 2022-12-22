package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kwantz/ecommerce/services/account/internal/entity"
	"github.com/kwantz/ecommerce/services/account/internal/utility"
)

type AccountController struct {
	accountUsecase AccountUsecase
}

func NewAccountController(accountUsecase AccountUsecase) *AccountController {
	return &AccountController{
		accountUsecase: accountUsecase,
	}
}

func (controller *AccountController) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	operation := "AccountController.CreateAccountHandler"

	request := entity.AccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("[%s] failed decode request body, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	response, err := controller.accountUsecase.CreateAccount(r.Context(), request)
	if err != nil {
		log.Printf("[%s] failed create account from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

func (controller *AccountController) GetAllAccountHandler(w http.ResponseWriter, r *http.Request) {
	operation := "AccountController.GetAllAccountHandler"

	response, err := controller.accountUsecase.GetAllAccount(r.Context())
	if err != nil {
		log.Printf("[%s] failed get all account from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

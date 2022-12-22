package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/kwantz/ecommerce/services/account/internal/entity"
	"github.com/kwantz/ecommerce/services/account/internal/utility"
)

type AuthController struct {
	authUsecase AuthUsecase
}

func NewAuthController(authUsecase AuthUsecase) *AuthController {
	return &AuthController{
		authUsecase: authUsecase,
	}
}

func (controller *AuthController) AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	operation := "AuthController.AuthenticationHandler"

	request := entity.LoginAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("[%s] failed decode request body, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	response, err := controller.authUsecase.Authentication(r.Context(), request)
	if err != nil {
		log.Printf("[%s] failed login account from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

func (controller *AuthController) AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	operation := "AuthController.AuthenticationHandler"

	authorizationHeader := r.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		log.Printf("[%s] invalid authorization header", operation)
		utility.ResponseErrorJSON(w, errors.New("bad request"))
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	response, err := controller.authUsecase.Authorization(r.Context(), tokenString)
	if err != nil {
		log.Printf("[%s] failed login account from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

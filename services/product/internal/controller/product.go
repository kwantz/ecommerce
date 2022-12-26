package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
	"github.com/kwantz/ecommerce/services/product/internal/utility"
)

type ProductController struct {
	productUsecase ProductUsecase
}

func NewProductController(productUsecase ProductUsecase) *ProductController {
	return &ProductController{
		productUsecase: productUsecase,
	}
}

func (controller *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	operation := "ProductController.CreateProductHandler"

	request := entity.ProductRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("[%s] failed decode request body, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	response, err := controller.productUsecase.CreateProduct(r.Context(), request)
	if err != nil {
		log.Printf("[%s] failed create product from usecase, cause: %s", operation, err.Error())
		utility.ResponseErrorJSON(w, err)
		return
	}

	utility.ResponseJSON(w, response)
}

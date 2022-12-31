package utility

import (
	"encoding/json"
	"net/http"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		ResponseErrorJSON(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByte)
}

func ResponseErrorJSON(w http.ResponseWriter, err error) {
	jsonByte, _ := json.Marshal(entity.ErrorResponse{
		Error: err.Error(),
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByte)
}

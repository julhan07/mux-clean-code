package utils

import (
	"encoding/json"
	"mux-clean-code/entity"
	"net/http"
)

func CreateErrorResponse(w http.ResponseWriter, msg string, err error) {

	res := &entity.ResponseApi{
		Error:   err,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func CreateSuccessResponse(w http.ResponseWriter, data *interface{}, msg string) {

	res := &entity.ResponseApi{
		Message: msg,
		Data:    data,
	}

	json.NewEncoder(w).Encode(res)
}

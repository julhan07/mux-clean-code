package presenter

import (
	"errors"
	"mux-clean-code/entity"
	"mux-clean-code/utils"
	"net/http"
)

func (u *presenter) GetUsers(w http.ResponseWriter, r *http.Request) {

	str := u.usecase.GetUsers(w, r)

	if str == "" {
		utils.CreateErrorResponse(w, entity.ERROR_RECEIVED, errors.New("error bro"))
	}

	utils.CreateSuccessResponse(w, nil, entity.DATA_RECEIVED)
}

func (u *presenter) GetUser(w http.ResponseWriter, r *http.Request) {
	str := u.usecase.GetUsers(w, r)

	if str == "" {
		utils.CreateErrorResponse(w, entity.ERROR_RECEIVED, errors.New("error bro"))
	}

	utils.CreateSuccessResponse(w, nil, entity.DATA_RECEIVED)
}

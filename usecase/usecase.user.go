package usecase

import "net/http"

func (u *usecase) GetUsers(w http.ResponseWriter, r *http.Request) string {
	return "All User"
}

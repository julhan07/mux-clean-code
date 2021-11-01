package repository

import (
	"net/http"
)

func (*repoGorm) GetUsers(w http.ResponseWriter, r *http.Request) string {
	return "All User"
}

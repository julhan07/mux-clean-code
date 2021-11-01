package usecase

import (
	"mux-clean-code/repository"
	"net/http"
)

type usecase struct {
	repo repository.RepositoryIndex
}

type UseCaseIndex interface {
	GetUsers(w http.ResponseWriter, r *http.Request) string
}

func NewUseCaseIndex(repo repository.RepositoryIndex) UseCaseIndex {
	return &usecase{
		repo: repo,
	}
}

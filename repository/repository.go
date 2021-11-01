package repository

import (
	"net/http"

	"gorm.io/gorm"
)

type repoGorm struct {
	db *gorm.DB
}

type RepositoryIndex interface {
	GetUsers(w http.ResponseWriter, r *http.Request) string
}

func NewRepositoryIndex(db *gorm.DB) RepositoryIndex {
	return &repoGorm{
		db: db,
	}
}

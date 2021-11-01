package main

import (
	"fmt"
	"log"
	"mux-clean-code/middleware"
	"mux-clean-code/presenter"
	"mux-clean-code/repository"
	"mux-clean-code/usecase"
	"net/http"
	"time"

	"mux-clean-code/database"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	db := database.IniDB()

	r.Use(middleware.LoggingMiddleware)

	repo := repository.NewRepositoryIndex(db)
	useCase := usecase.NewUseCaseIndex(repo)

	app := presenter.NewPresenterIndex(useCase, r)

	srv := &http.Server{
		Handler: app,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("App running on PORT 8000")

	log.Fatal(srv.ListenAndServe())
}

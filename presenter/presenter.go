package presenter

import (
	"mux-clean-code/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type presenter struct {
	usecase usecase.UseCaseIndex
	*mux.Router
}

func NewPresenterIndex(usecase usecase.UseCaseIndex, r *mux.Router) http.Handler {
	p := &presenter{
		usecase: usecase,
		Router:  r,
	}

	p.routes()
	return *&p
}

func (p *presenter) routes() {
	p.HandleFunc("/user", p.GetUsers).Methods("GET")
	p.HandleFunc("/user/{id}", p.GetUser).Methods("GET")
}

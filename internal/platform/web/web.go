package web

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Handler func(http.ResponseWriter, *http.Request) error

type App struct {
	mux *chi.Mux
	log *log.Logger
}

func NewApp(logger *log.Logger) *App {
	return &App{
		mux: chi.NewRouter(),
		log: logger,
	}
}

func (a *App) Handle(method, pattern string, h Handler) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			a.log.Printf("Error: %v", err)
			if err := RespondError(w, err); err != nil {
				a.log.Printf("Error: %v", err)
			}
		}
	}

	a.mux.MethodFunc(method, pattern, fn)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

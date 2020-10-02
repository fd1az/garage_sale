package handlers

import (
	"log"
	"net/http"

	"github.com/fdiaz7/garage_sales/internal/platform/web"
	"github.com/jmoiron/sqlx"
)

func API(logger *log.Logger, db *sqlx.DB) http.Handler {
	p := Products{DB: db, Log: logger}
	app := web.NewApp(logger)
	app.Handle(http.MethodGet, "/v1/products", p.List)
	app.Handle(http.MethodPost, "/v1/products", p.Create)
	app.Handle(http.MethodGet, "/v1/products/{id}", p.Retrieve)

	return app
}

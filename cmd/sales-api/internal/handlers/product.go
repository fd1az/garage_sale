package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"github.com/fdiaz7/garage_sales/internal/platform/web"
	"github.com/fdiaz7/garage_sales/internal/product"
	"github.com/jmoiron/sqlx"
)

type Products struct {
	DB  *sqlx.DB
	Log *log.Logger
}

// ListProducts show products
func (p *Products) List(w http.ResponseWriter, r *http.Request) error {
	//Empty value list no a zero value list- esto facilita a los consumidores de la api no recibir un null sino un []
	list, err := product.List(p.DB)

	if err != nil {
		return err
	}

	return web.Respond(w, list, http.StatusOK)

}

// Retrieve a single product
func (p *Products) Retrieve(w http.ResponseWriter, r *http.Request) error {
	//Empty value list no a zero value list- esto facilita a los consumidores de la api no recibir un null sino un []
	id := chi.URLParam(r, "id")
	prod, err := product.Retrieve(p.DB, id)

	if err != nil {
		switch err {
		case product.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		case product.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		default:
			return errors.Wrapf(err, "looking for product %q", id)
		}

	}

	return web.Respond(w, prod, http.StatusOK)

}

//Create decode a JSON document from post request and create a new product
func (p *Products) Create(w http.ResponseWriter, r *http.Request) error {
	var np product.NewProduct
	if err := web.Decoder(r, &np); err != nil {
		return err
	}

	prod, err := product.Create(p.DB, np, time.Now())

	if err != nil {
		return err
	}

	return web.Respond(w, prod, http.StatusCreated)
}

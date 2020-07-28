package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fdiaz7/garage_sales/internal/product"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	DB *sqlx.DB
}

// ListProducts show products
func (p *Product) List(w http.ResponseWriter, r *http.Request) {
	//Empty value list no a zero value list- esto facilita a los consumidores de la api no recibir un null sino un []
	list, err := product.List(p.DB)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error querying db", err)
	}

	data, err := json.Marshal(list)
	if err != nil {
		log.Println("Error writing", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		log.Println("Error writing", err)
	}

}

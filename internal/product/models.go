package product

import (
	"time"
)

//Product to sell
type Product struct {
	ID          string    `db:"product_id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Cost        int       `db:"cost" json:"cost"`
	Quantity    int       `db:"quantity" json:"quantity"`
	DateUpdated time.Time `db:"date_updated" json:"date_updated"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
}

//NewProduct is what require from clients to create products
type NewProduct struct {
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	Quantity int    `json:"quantity"`
}

package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Id    uuid.UUID `db:"product_id"`
	Name  string    `db:"name"`
	Price float64   `db:"price"`
	Values []CurrencyExchange
}

type AddProduct struct {
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}

type Currency struct {
	Code string `json:"code"`
	Value string `json:"high"`
}

type CurrencyExchange struct {
	Code string `json:"code"`
	Value float64 `json:"value"`
}
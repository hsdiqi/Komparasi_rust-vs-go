package model

import "time"

type Product struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	Stock       int       `db:"stock" json:"stock"`
	Category    string    `db:"category" json:"category"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gte=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
	Category    string  `json:"category" binding:"required"`
}

type ProductListResponse struct {
	Status string    `json:"status"`
	Data   []Product `json:"data"`
	Meta   Meta      `json:"meta"`
}

type ProductResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

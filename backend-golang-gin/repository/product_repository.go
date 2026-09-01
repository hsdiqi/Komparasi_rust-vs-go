package repository

import (
	"context"

	"backend-golang-gin/model"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct { db *sqlx.DB }
func NewProductRepository(db *sqlx.DB) *ProductRepository { return &ProductRepository{db: db} }

func (r *ProductRepository) FindAll(ctx context.Context, page, limit int) ([]model.Product, error) {
	offset := (page - 1) * limit
	products := []model.Product{}
	err := r.db.SelectContext(ctx, &products, `SELECT id, name, description, price, stock, category, created_at, updated_at FROM products ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
	return products, err
}

func (r *ProductRepository) FindByID(ctx context.Context, id int64) (model.Product, error) {
	var p model.Product
	err := r.db.GetContext(ctx, &p, `SELECT id, name, description, price, stock, category, created_at, updated_at FROM products WHERE id=$1`, id)
	return p, err
}

func (r *ProductRepository) Search(ctx context.Context, keyword string, page, limit int) ([]model.Product, error) {
	offset := (page - 1) * limit
	products := []model.Product{}
	pattern := "%" + keyword + "%"
	err := r.db.SelectContext(ctx, &products, `SELECT id, name, description, price, stock, category, created_at, updated_at FROM products WHERE name ILIKE $1 OR category ILIKE $1 ORDER BY id LIMIT $2 OFFSET $3`, pattern, limit, offset)
	return products, err
}

func (r *ProductRepository) Create(ctx context.Context, req model.ProductRequest) (model.Product, error) {
	var p model.Product
	err := r.db.GetContext(ctx, &p, `INSERT INTO products (name, description, price, stock, category) VALUES ($1,$2,$3,$4,$5) RETURNING id, name, description, price, stock, category, created_at, updated_at`, req.Name, req.Description, req.Price, req.Stock, req.Category)
	return p, err
}

func (r *ProductRepository) Update(ctx context.Context, id int64, req model.ProductRequest) (model.Product, error) {
	var p model.Product
	err := r.db.GetContext(ctx, &p, `UPDATE products SET name=$1, description=$2, price=$3, stock=$4, category=$5, updated_at=NOW() WHERE id=$6 RETURNING id, name, description, price, stock, category, created_at, updated_at`, req.Name, req.Description, req.Price, req.Stock, req.Category, id)
	return p, err
}

func (r *ProductRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM products WHERE id=$1`, id)
	return err
}

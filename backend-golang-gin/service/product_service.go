package service

import (
	"context"

	"backend-golang-gin/model"
	"backend-golang-gin/repository"
)

type ProductService struct { repo *repository.ProductRepository }
func NewProductService(repo *repository.ProductRepository) *ProductService { return &ProductService{repo: repo} }

func (s *ProductService) List(ctx context.Context, page, limit int) ([]model.Product, error) { return s.repo.FindAll(ctx, page, limit) }
func (s *ProductService) Detail(ctx context.Context, id int64) (model.Product, error) { return s.repo.FindByID(ctx, id) }
func (s *ProductService) Search(ctx context.Context, keyword string, page, limit int) ([]model.Product, error) { return s.repo.Search(ctx, keyword, page, limit) }
func (s *ProductService) Create(ctx context.Context, req model.ProductRequest) (model.Product, error) { return s.repo.Create(ctx, req) }
func (s *ProductService) Update(ctx context.Context, id int64, req model.ProductRequest) (model.Product, error) { return s.repo.Update(ctx, id, req) }
func (s *ProductService) Delete(ctx context.Context, id int64) error { return s.repo.Delete(ctx, id) }

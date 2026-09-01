use crate::{model::{Product, ProductRequest}, repository::ProductRepository};

#[derive(Clone)]
pub struct ProductService { repo: ProductRepository }

impl ProductService {
    pub fn new(repo: ProductRepository) -> Self { Self { repo } }
    pub async fn list(&self, page: i64, limit: i64) -> Result<Vec<Product>, sqlx::Error> { self.repo.find_all(page, limit).await }
    pub async fn detail(&self, id: i64) -> Result<Product, sqlx::Error> { self.repo.find_by_id(id).await }
    pub async fn search(&self, keyword: String, page: i64, limit: i64) -> Result<Vec<Product>, sqlx::Error> { self.repo.search(keyword, page, limit).await }
    pub async fn create(&self, req: ProductRequest) -> Result<Product, sqlx::Error> { self.repo.create(req).await }
    pub async fn update(&self, id: i64, req: ProductRequest) -> Result<Product, sqlx::Error> { self.repo.update(id, req).await }
    pub async fn delete(&self, id: i64) -> Result<u64, sqlx::Error> { self.repo.delete(id).await }
}

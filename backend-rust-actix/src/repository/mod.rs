use sqlx::PgPool;
use crate::model::{Product, ProductRequest};

#[derive(Clone)]
pub struct ProductRepository { pool: PgPool }

impl ProductRepository {
    pub fn new(pool: PgPool) -> Self { Self { pool } }

    pub async fn find_all(&self, page: i64, limit: i64) -> Result<Vec<Product>, sqlx::Error> {
        let offset = (page - 1) * limit;
        sqlx::query_as::<_, Product>("SELECT id, name, description, price, stock, category, created_at, updated_at FROM products ORDER BY id LIMIT $1 OFFSET $2")
            .bind(limit)
            .bind(offset)
            .fetch_all(&self.pool)
            .await
    }

    pub async fn find_by_id(&self, id: i64) -> Result<Product, sqlx::Error> {
        sqlx::query_as::<_, Product>("SELECT id, name, description, price, stock, category, created_at, updated_at FROM products WHERE id=$1")
            .bind(id)
            .fetch_one(&self.pool)
            .await
    }

    pub async fn search(&self, keyword: String, page: i64, limit: i64) -> Result<Vec<Product>, sqlx::Error> {
        let offset = (page - 1) * limit;
        let pattern = format!("%{}%", keyword);
        sqlx::query_as::<_, Product>("SELECT id, name, description, price, stock, category, created_at, updated_at FROM products WHERE name ILIKE $1 OR category ILIKE $1 ORDER BY id LIMIT $2 OFFSET $3")
            .bind(pattern)
            .bind(limit)
            .bind(offset)
            .fetch_all(&self.pool)
            .await
    }

    pub async fn create(&self, req: ProductRequest) -> Result<Product, sqlx::Error> {
        sqlx::query_as::<_, Product>("INSERT INTO products (name, description, price, stock, category) VALUES ($1,$2,$3,$4,$5) RETURNING id, name, description, price, stock, category, created_at, updated_at")
            .bind(req.name)
            .bind(req.description)
            .bind(req.price)
            .bind(req.stock)
            .bind(req.category)
            .fetch_one(&self.pool)
            .await
    }

    pub async fn update(&self, id: i64, req: ProductRequest) -> Result<Product, sqlx::Error> {
        sqlx::query_as::<_, Product>("UPDATE products SET name=$1, description=$2, price=$3, stock=$4, category=$5, updated_at=NOW() WHERE id=$6 RETURNING id, name, description, price, stock, category, created_at, updated_at")
            .bind(req.name)
            .bind(req.description)
            .bind(req.price)
            .bind(req.stock)
            .bind(req.category)
            .bind(id)
            .fetch_one(&self.pool)
            .await
    }

    pub async fn delete(&self, id: i64) -> Result<u64, sqlx::Error> {
        let result = sqlx::query("DELETE FROM products WHERE id=$1")
            .bind(id)
            .execute(&self.pool)
            .await?;
        Ok(result.rows_affected())
    }
}

use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use sqlx::FromRow;

#[derive(Debug, Serialize, FromRow)]
pub struct Product {
    pub id: i64,
    pub name: String,
    pub description: String,
    pub price: f64,
    pub stock: i32,
    pub category: String,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}

#[derive(Debug, Deserialize)]
pub struct ProductRequest {
    pub name: String,
    pub description: String,
    pub price: f64,
    pub stock: i32,
    pub category: String,
}

#[derive(Debug, Deserialize)]
pub struct PaginationQuery {
    pub page: Option<i64>,
    pub limit: Option<i64>,
}

#[derive(Debug, Deserialize)]
pub struct SearchQuery {
    pub keyword: Option<String>,
    pub page: Option<i64>,
    pub limit: Option<i64>,
}

#[derive(Debug, Serialize)]
pub struct Meta {
    pub page: i64,
    pub limit: i64,
}

#[derive(Debug, Serialize)]
pub struct ListResponse<T> {
    pub status: String,
    pub data: Vec<T>,
    pub meta: Meta,
}

#[derive(Debug, Serialize)]
pub struct SingleResponse<T: Serialize> {
    pub status: String,
    pub data: T,
}

#[derive(Debug, Serialize)]
pub struct DeleteResponse {
    pub status: String,
    pub data: Option<serde_json::Value>,
}

#[derive(Debug, Serialize)]
pub struct ErrorResponse {
    pub status: String,
    pub message: String,
}

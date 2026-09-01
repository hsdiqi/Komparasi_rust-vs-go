use actix_web::{web, HttpResponse, Responder};
use sqlx::Error;
use crate::model::{DeleteResponse, ErrorResponse, ListResponse, Meta, PaginationQuery, ProductRequest, SearchQuery, SingleResponse};
use crate::service::ProductService;

fn normalize_page_limit(page: Option<i64>, limit: Option<i64>) -> (i64, i64) {
    let p = page.unwrap_or(1).max(1);
    let l = limit.unwrap_or(10).clamp(1, 100);
    (p, l)
}

fn db_error(err: Error) -> HttpResponse {
    match err {
        Error::RowNotFound => HttpResponse::NotFound().json(ErrorResponse { status: "error".to_string(), message: "product not found".to_string() }),
        _ => HttpResponse::InternalServerError().json(ErrorResponse { status: "error".to_string(), message: err.to_string() }),
    }
}

pub async fn list(service: web::Data<ProductService>, query: web::Query<PaginationQuery>) -> impl Responder {
    let (page, limit) = normalize_page_limit(query.page, query.limit);
    match service.list(page, limit).await {
        Ok(data) => HttpResponse::Ok().json(ListResponse { status: "success".to_string(), data, meta: Meta { page, limit } }),
        Err(err) => db_error(err),
    }
}

pub async fn detail(service: web::Data<ProductService>, path: web::Path<i64>) -> impl Responder {
    match service.detail(path.into_inner()).await {
        Ok(data) => HttpResponse::Ok().json(SingleResponse { status: "success".to_string(), data }),
        Err(err) => db_error(err),
    }
}

pub async fn search(service: web::Data<ProductService>, query: web::Query<SearchQuery>) -> impl Responder {
    let (page, limit) = normalize_page_limit(query.page, query.limit);
    let keyword = query.keyword.clone().unwrap_or_default();
    match service.search(keyword, page, limit).await {
        Ok(data) => HttpResponse::Ok().json(ListResponse { status: "success".to_string(), data, meta: Meta { page, limit } }),
        Err(err) => db_error(err),
    }
}

pub async fn create(service: web::Data<ProductService>, body: web::Json<ProductRequest>) -> impl Responder {
    match service.create(body.into_inner()).await {
        Ok(data) => HttpResponse::Created().json(SingleResponse { status: "success".to_string(), data }),
        Err(err) => db_error(err),
    }
}

pub async fn update(service: web::Data<ProductService>, path: web::Path<i64>, body: web::Json<ProductRequest>) -> impl Responder {
    match service.update(path.into_inner(), body.into_inner()).await {
        Ok(data) => HttpResponse::Ok().json(SingleResponse { status: "success".to_string(), data }),
        Err(err) => db_error(err),
    }
}

pub async fn delete(service: web::Data<ProductService>, path: web::Path<i64>) -> impl Responder {
    match service.delete(path.into_inner()).await {
        Ok(_) => HttpResponse::Ok().json(DeleteResponse { status: "success".to_string(), data: None }),
        Err(err) => db_error(err),
    }
}

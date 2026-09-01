use actix_web::{web, HttpResponse};
use crate::handler;

pub fn configure(cfg: &mut web::ServiceConfig) {
    cfg.route("/health", web::get().to(|| async { HttpResponse::Ok().json(serde_json::json!({"status":"ok"})) }))
        .route("/products", web::get().to(handler::list))
        .route("/products/search", web::get().to(handler::search))
        .route("/products/{id}", web::get().to(handler::detail))
        .route("/products", web::post().to(handler::create))
        .route("/products/{id}", web::put().to(handler::update))
        .route("/products/{id}", web::delete().to(handler::delete));
}

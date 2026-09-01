mod config;
mod db;
mod model;
mod repository;
mod service;
mod handler;
mod route;

use actix_web::{web, App, HttpServer};
use config::Config;
use repository::ProductRepository;
use service::ProductService;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let cfg = Config::from_env();
    let pool = db::new_pool(&cfg).await.expect("failed to connect database");
    let repo = ProductRepository::new(pool);
    let service = ProductService::new(repo);
    let addr = format!("0.0.0.0:{}", cfg.app_port);
    println!("Rust Actix Web running on {}", addr);

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(service.clone()))
            .configure(route::configure)
    })
    .bind(addr)?
    .run()
    .await
}

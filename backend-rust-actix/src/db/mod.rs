use sqlx::{postgres::PgPoolOptions, PgPool};
use crate::config::Config;

pub async fn new_pool(cfg: &Config) -> Result<PgPool, sqlx::Error> {
    PgPoolOptions::new()
        .max_connections(cfg.max_connections)
        .connect(&cfg.database_url)
        .await
}

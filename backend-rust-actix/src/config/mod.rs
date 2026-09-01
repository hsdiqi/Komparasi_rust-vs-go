use std::env;

#[derive(Clone)]
pub struct Config {
    pub app_port: u16,
    pub database_url: String,
    pub max_connections: u32,
}

impl Config {
    pub fn from_env() -> Self {
        dotenvy::dotenv().ok();
        Self {
            app_port: env::var("APP_PORT").unwrap_or_else(|_| "8081".to_string()).parse().unwrap_or(8081),
            database_url: env::var("DATABASE_URL").unwrap_or_else(|_| "postgres://postgres:postgres@localhost:5432/perf_db?sslmode=disable".to_string()),
            max_connections: env::var("DB_MAX_CONNECTIONS").unwrap_or_else(|_| "25".to_string()).parse().unwrap_or(25),
        }
    }
}

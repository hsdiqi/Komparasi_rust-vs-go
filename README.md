# Backend Performance Research: Golang Gin vs Rust Actix Web

Project eksperimen REST API CRUD produk dengan PostgreSQL untuk penelitian skripsi/jurnal.

## Services
- Golang Gin: http://localhost:8080
- Rust Actix Web: http://localhost:8081
- PostgreSQL: localhost:5432

## Run
```bash
docker compose up --build
```

## Seed dataset
```bash
docker exec -i perf-postgres psql postgres://postgres:postgres@localhost:5432/perf_db < backend-golang-gin/migrations/001_create_products.sql
sed 's/:COUNT/10000/g' backend-golang-gin/scripts/seed.sql | docker exec -i perf-postgres psql -U postgres -d perf_db
sed 's/:COUNT/50000/g' backend-golang-gin/scripts/seed.sql | docker exec -i perf-postgres psql -U postgres -d perf_db
sed 's/:COUNT/100000/g' backend-golang-gin/scripts/seed.sql | docker exec -i perf-postgres psql -U postgres -d perf_db
```

## Health check
```bash
curl http://localhost:8080/health
curl http://localhost:8081/health
```

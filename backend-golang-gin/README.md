# backend-golang-gin

REST API CRUD produk untuk eksperimen performa Gin vs Actix Web.

## Run lokal
```bash
cp .env.example .env
go mod tidy
go run ./cmd/server
```

## Migrasi
```bash
psql "$DATABASE_URL" -f migrations/001_create_products.sql
```

## Seeder
```bash
./scripts/seed.sh 10000
./scripts/seed.sh 50000
./scripts/seed.sh 100000
```

## Endpoint
- GET /products?page=1&limit=10
- GET /products/:id
- GET /products/search?keyword=phone
- POST /products
- PUT /products/:id
- DELETE /products/:id

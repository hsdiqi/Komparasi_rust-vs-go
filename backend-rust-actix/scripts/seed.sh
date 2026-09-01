#!/usr/bin/env bash
set -euo pipefail
COUNT=${1:-10000}
DB_URL=${DATABASE_URL:-postgres://postgres:postgres@localhost:5432/perf_db?sslmode=disable}
sed "s/:COUNT/$COUNT/g" scripts/seed.sql | psql "$DB_URL"

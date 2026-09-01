# Benchmark Commands

## k6
```bash
BASE_URL=http://localhost:8080 VUS=10 DURATION=30s k6 run k6-read.js
BASE_URL=http://localhost:8081 VUS=10 DURATION=30s k6 run k6-read.js
```

## autocannon
```bash
autocannon -c 10 -a 1000 http://localhost:8080/products?page=1\&limit=10
autocannon -c 10 -a 1000 http://localhost:8081/products?page=1\&limit=10
```

## wrk
```bash
wrk -t4 -c10 -d30s http://localhost:8080/products?page=1\&limit=10
wrk -t4 -c10 -d30s http://localhost:8081/products?page=1\&limit=10
```

## Apache Benchmark
```bash
ab -n 1000 -c 10 http://localhost:8080/products?page=1\&limit=10
ab -n 1000 -c 10 http://localhost:8081/products?page=1\&limit=10
```

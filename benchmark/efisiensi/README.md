# cara ambil pid container 
- go
```bash
docker inspect -f '{{.State.Pid}}' backend-golang-gin
```
- rust
```bash
docker inspect -f '{{.State.Pid}}' backend-rust-actix
```

# Cara run program 

```bash
sudo python3 monitor_resources.py PID hasil.csv
```
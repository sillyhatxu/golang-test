CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
docker build -t xushikuan/docker-replicas .
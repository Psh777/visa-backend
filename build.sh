GOOS=linux GOARCH=amd64 go build -o main main.go
docker build -t psypsy777/visa-backend -f Dockerfile .
docker push psypsy777/visa-backend:latest

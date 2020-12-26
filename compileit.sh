env GOOS=linux GOARCH=amd64 go build -o server -v cmd/server/main.go
# docker build -t sample .
# docker run -p3333:3333 sample

GOOS=linux GOARCH=amd64 go build  -a -tags netgo -ldflags '-w -s'
docker build --tag linehaul .
docker run --rm -p 8080:8080 linehaul
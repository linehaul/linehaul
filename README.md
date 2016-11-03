# linehaul cd server

## Development Guide

### Building
go build -o release/linehaul linehaul

### Updating GoDeps
To update the go dependencies, use:

    ````
        go get -u github.com/kardianos/govendor
        govendor add +external
    ````
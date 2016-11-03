# linehaul cd server

## Development Guide

### Building
````
    ./build.sh
````

### Updating GoDeps
To update the go dependencies, use:

    ````
        go get -u github.com/kardianos/govendor
        govendor add +external
    ````
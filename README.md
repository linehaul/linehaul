# linehaul cd server

## Development Guide

### Prerequisites
You require:

 - Docker must be installed on and on $PATH
 - golang 1.7
 - govendor
    - Install with:
        ```
        go get -u github.com/kardianos/govendor
        ```

### Building
To build locally, run the following:

```
govendor sync
./build.sh
```

Then simply navigate to http://localhost:8080/

### Updating GoDeps
To update the go dependencies, use:

```
go get -u github.com/kardianos/govendor
govendor add +external
```
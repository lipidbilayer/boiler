# Backend

## Installation
Requirement: 
- Go 1.19 with env GOPATH  
  
  
    
### CLI Tools Installation
#### Install migrate CLI
```bash
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
```
  
#### Install statik
```bash
$ go install github.com/rakyll/statik@v0.1.7
```
  
#### Install Task
```bash
$ go install github.com/go-task/task/v3/cmd/task@latest
```
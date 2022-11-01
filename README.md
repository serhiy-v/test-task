# REST API for Test Task

Service for uploading, filtering and gettig information about transactions

## Starting Up

**1. Clone repository**

```bash

  # clone branch
  git clone git@github.com:serhiy-v/test-task.git

```

**2.Create DB and API Doc server**

```bash
   cd test-api
   docker-compose up -d #To run in detached mode
```
API documentation [here](http://localhost:8080)

**3.Start the server**

```bash
   go run cmd/main.go 
```
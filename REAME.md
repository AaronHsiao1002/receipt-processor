# Receipt Processor

## Overview

This is a web service that processes receipts and awards points based on specific rules.

## Running the Application

### Locally

1. Ensure Go is installed.
2. Install all dependencies.

   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

### Run the app using Docker

```bash
docker build -t receipt-processor .
```

```bash
docker run -p 8080:8080 receipt-processor
```

## Testing the Application

### 1. Use Curl or Postman to send requests to `localhost:8080`

or

### 2. Run unit tests

```bash
go test ./...
```

# Financial Tracker

Financial Tracker is a Go language web application that allows you to track your income and expenses.

## Functionality

- Adding transactions (income and expenses) with amount, currency, type (income or expense) and category.
- Obtaining a current balance that takes into account all income and expenses.

## Project structure

- `api/handlers.go` - handlers for API requests.
- `api/router.go` - route configuration for API.
- `db/db.go` - database initialization.
- `db/migrations/schema.sql` - database schema.
- `models/category.go` - model for transaction categories.
- `models/transaction.go` - model for transactions.
- `main.go` - the main entry point for the application.
- `docker-compose.yml` - Docker Compose settings for running the application and database containers.
- `Dockerfile` - instructions for building a Docker image of the application.

## Requirements

- Docker
- Docker Compose

### Install and run

### Cloning the repository
```bash
git clone https://github.com/DebugSensei/Financial-Tracker.git
```
### Run
```bash
cd Financial-Tracker
docker-compose up --build
```
### Functionality check
To add a transaction, send a POST request to /transaction with the request body in JSON format:
```bash
{
    “date”: “2023-05-01T12:00:00Z”,
    “amount”: 100,
    “currency”: “USD”,
    “type”: “income”,
    “category_id”: 1
}
```
### Example of use with CMD (Income):
```bash
curl -X POST "http://localhost:8080/transaction" -H "Content-Type: application/json" -d "{ \"date\": \"2024-05-01T12:19:18Z\", \"amount\": 100, \"currency\": \"USD\", \"type\": \"income\", \"category_id\": 1 }"
```
### Example of use with CMD (Expense):
```bash
curl -X POST "http://localhost:8080/transaction" -H "Content-Type: application/json" -d "{ \"date\": \"2023-05-01T12:00:00Z\", \"amount\": 50.00, \"currency\": \"USD\", \"type\": \"expense\", \"category_id\": 2 }"
```
### Receive Balance
To get the current balance, send a GET request to /balance.
CMD example(check balance):
```bash
curl -X GET "http://localhost:8080/balance"
```
The response will be in json format:
```bash
{“balance”: 201.00}
```

### Example of use with PowerShell (Income):
```bash
$headers = @{
    “Content-Type” = “application/json”
}

$body = @{
    date = “2023-05-01T12:00:00Z”
    amount = 100
    currency = “USD”
    type = “income”
    category_id = 1
} | ConvertTo-Json

Invoke-WebRequest -Uri “http://localhost:8080/transaction” -Method POST -Headers $headers -Body $body
```
### Example of use with PowerShell (Expense):
```bash
$headers = @{
    "Content-Type" = "application/json"
}

$body = @{
    date = "2023-05-01T12:00:00Z"
    amount = 50.00
    currency = "USD"
    type = "expense"
    category_id = 2
} | ConvertTo-Json

Invoke-WebRequest -Uri "http://localhost:8080/transaction" -Method POST -Headers $headers -Body $body
```

### Receive Balance
To get the current balance, send a GET request to /balance.
PowerShell example:
```bash
Invoke-WebRequest -Uri “http://localhost:8080/balance” -Method GET
```
The response will be in json format:
```bash
{“balance”: 201.00}
```


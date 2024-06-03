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
    “amount”: 100.50,
    “currency”: “USD”,
    “type”: “income”,
    “category_id”: 1
}
```

### Example of use with PowerShell:
```bash
$headers = @{
    “Content-Type” = “application/json”
}

$body = @{
    date = “2023-05-01T12:00:00Z”
    amount = 100.50
    currency = “USD”
    type = “income”
    category_id = 1
} | ConvertTo-Json

Invoke-WebRequest -Uri “http://localhost:8080/transaction” -Method POST -Headers $headers -Body $body
```

### Receive Balance
To get the current balance, send a GET request to /balance.
PowerShell example:

```bash
Invoke-WebRequest -Uri “http://localhost:8080/balance” -Method GET
```
The response will be in json format:
```bash
{
    “balance”: 201.00
}

```

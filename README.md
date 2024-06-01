# Financial Tracker

Financial Tracker - это веб-приложение на языке Go, которое позволяет отслеживать доходы и расходы.

## Функциональность

- Добавление транзакций (доходы и расходы) с указанием суммы, валюты, типа (доход или расход) и категории.
- Получение текущего баланса, учитывающего все доходы и расходы.

## Структура проекта

- `api/handlers.go` - обработчики для API запросов.
- `api/router.go` - настройка маршрутов для API.
- `db/db.go` - инициализация базы данных.
- `db/migrations/schema.sql` - схема базы данных.
- `models/category.go` - модель для категорий транзакций.
- `models/transaction.go` - модель для транзакций.
- `main.go` - основная точка входа для приложения.
- `docker-compose.yml` - настройки Docker Compose для запуска контейнеров приложения и базы данных.
- `Dockerfile` - инструкции для сборки Docker-образа приложения.

## Требования

- Docker
- Docker Compose

## Установка и запуск

### Клонирование репозитория
```bash
git clone https://github.com/DebugSensei/Financial-Tracker.git
```
### Запуск
```bash
cd Financial-Tracker
docker-compose up --build
```
### Проверка работоспособности
Для добавления транзакции отправьте POST-запрос на /transaction с телом запроса в формате JSON:
```bash
{
    "date": "2023-05-01T12:00:00Z",
    "amount": 100.50,
    "currency": "USD",
    "type": "income",
    "category_id": 1
}
```

Пример использования c использованием PowerShell:
```bash
$headers = @{
    "Content-Type" = "application/json"
}

$body = @{
    date = "2023-05-01T12:00:00Z"
    amount = 100.50
    currency = "USD"
    type = "income"
    category_id = 1
} | ConvertTo-Json

Invoke-WebRequest -Uri "http://localhost:8080/transaction" -Method POST -Headers $headers -Body $body
```

### Получение баланса
Для получения текущего баланса отправьте GET-запрос на /balance.
Пример использования c использованием PowerShell:

```bash
Invoke-WebRequest -Uri "http://localhost:8080/balance" -Method GET
```
Ответ будет в формате json:
```json
{
    "balance": 201.00
}

```



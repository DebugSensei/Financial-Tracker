FROM golang:1.18-alpine

# Установите PostgreSQL клиент
RUN apk update && apk add postgresql-client

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD ["sh", "-c", "PGPASSWORD=$DB_PASSWORD psql -h db -U $DB_USER -d $DB_NAME -f db/migrations/schema.sql && ./main"]

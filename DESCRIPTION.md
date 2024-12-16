# Перед первым запуском
```sh
cp .env.sample .env
```

# Запуск
```sh
go run ./cmd/web/main.go
```

# Пример работы API

## Заказ успешно создался

```sh
curl --location --request POST 'localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "hotel_id": 1,
    "room_id": 2,
    "email": "test@test.ru",
    "from": "2025-01-02T00:00:00Z",
    "to": "2025-01-03T00:00:00Z"
}'
```

## Ответ

```
201
{
    "id": 1,
    "price": 5000,
    "status": "ready_to_pay",
    "email": "test@test.ru",
    "bookings": [
        {
            "room_id": 2,
            "from": "2025-01-02T00:00:00Z",
            "to": "2025-01-03T00:00:00Z"
        }
    ]
}
```

## Ответ (ошибка) валидация
```
422
{
    "code": "Unprocessable Entity",
    "message": "input validation errors",
    "error": [
        "room_id is required and must be > 0",
        "hotel_id is required and must be > 0",
        "email is required",
        "from and to is required, from less than to, from and to more than now"
    ]
}
```

## Ответ (ошибка) нет свободных дат
```
422
{
    "code": "Unprocessable Entity",
    "message": "room not available for selected dates"
}
```


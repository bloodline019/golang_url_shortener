# golang_url_shortener

# Использование

API реализует следующие методы:
1. POST /create_short_url - принимает JSON с длинной ссылкой и id пользователя, возвращает укороченную ссылку
2. GET /:shortUrl - перенаправляет на исходную ссылку при передаче в метод укороченной ссылки
3. POST /qrcode - укорачивает ссылку и возвращает QR по которому доступно перенаправление на исходную ссылку

Формат работы с POST /create_short_url
```
curl -X POST http://localhost:8080/create_short_url \
-H "Content-Type: application/json" \
-d '{"long_url": "https://www.google.com", "user_id" : "wqer7c6-7wer-qnof-fneiun"}'
```

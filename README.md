# wcharge_back
## Запуск
## 
## Api
### UserApi
#### 1. Регистрация пользователя:
- Запрос
```http request
POST /api/user/create HTTP 1./1
Content-Type: application/json

{"username":"username", "password":"password","email":"user@mail.com" "address_id":1,"role_id":1}
```
- Ответ
```http request
HTTP/1.1 200 OK
```
#### 2. Login:
- Запрос
```http request
POST /api/user/login HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/118.0
Content-Type: application/json

{"username":"user1","password":"user1"}
```
- Ответ
```http request
HTTP/1.1 200 OK
Set-Cookie: userSession=MTY5NjE2NTE0OXxEWDhFQVFMX2dBQUJFQUVRQUFBZV80QUFBUVp6ZEhKcGJtY01DUUFIZFhObGNsOXBaQU5wYm5RRUFnQUN8_CBaLGbfy_xFp8L0wYwCvD_cYUhQtTvQM-z_OMPxG9Q=; Path=/; Expires=Tue, 31 Oct 2023 12:59:09 GMT; Max-Age=2592000
Date: Sun, 01 Oct 2023 12:59:09 GMT


```
#### 3. Получить всех пользователей:
- Запрос
```http request
GET /api/user/all HTTP/1.1
Host: 127.0.0.1:8080
```
- Ответ
```http request
HTTP/1.1 200 OK
Content-Type: application/json

...
```
#### 4. Получить пользователя по id:
- Запрос
```http request
GET /api/user/get/{id} HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/118.0
```
- Ответ
```http request
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 01 Oct 2023 13:07:50 GMT

...
```
#### 5. Обновить пользователя
- Запрос
```http request
PUT /api/user/update/2 HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/118.0
Content-Type: application/json

{"username":"user","email":"user@mail.com","role":1,"address_id":1}
```
- Ответ
```http request
HTTP/1.1 200 OK
```
#### 5. Удалить пользователя
- Запрос
```http request
DELETE /api/user/delete/1 HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/118.0
```
- Ответ
```http request
HTTP/1.1 200 OK
```
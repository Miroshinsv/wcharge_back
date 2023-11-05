# wcharge_back
## Запуск
1. Создать переменные окружения (```.env```)
2. ```go run ./cmd/app/main.go```
## Api
```
POST /v1/login
POST /v1/logout
GET /v1/api/whoami
-- PUT v1/api/user/{id:[0-9]+}/set-role

GET /v1/api/powerbank/all
GET /v1/api/powerbank/get/{id:[0-9]+}
POST /v1/api/powerbank/create
PUT /v1/api/powerbank/update/{id:[0-9]+}
DELETE /v1/api/powerbank/delete/{id:[0-9]+}

GET /v1/api/station/all
GET /v1/api/station/get/{id:[0-9]+}
POST /v1/api/station/create
PUT /v1/api/station/update/{id:[0-9]+}
DELETE /v1/api/station/delete/{id:[0-9]+}
GET /v1/api/station/{id:[0-9]+}/get/all-powerbanks
POST /v1/api/station/{station_id:[0-9]+}/take-powerbank/{powerbank_id:[0-9]+}
POST /v1/api/station/{station_id:[0-9]+}/put-powerbank/{powerbank_id:[0-9]+}
POST /v1/app/station/{station_id:[0-9]+}/add-powerbank/{powerbank_id:[0-9]+}

GET /v1/api/user/all
GET /v1/api/user/get/{id:[0-9]+}
POST /v1/api/user/create
PUT /v1/api/user/update/{id:[0-9]+}
DELETE /v1/api/user/delete/{id:[0-9]+}
GET /v1/api/user/{id:[0-9]+}/get/all-powerbanks

GET /v1/api/address/all
GET /v1/api/address/get/{id:[0-9]+}
POST /v1/api/address/create
PUT /v1/api/address/update/{id:[0-9]+}
DELETE /v1/api/address/delete/{id:[0-9]+}   
```
### UserApi
#### 1. Регистрация пользователя:
- Запрос
```http request
POST /v1/api/user/create HTTP 1./1
Content-Type: application/json

{"username":"username", "password":"password","email":"user@mail.com" "address_id":1,"role_id":1}
```
#### 2. Login:
- Запрос
```http request
POST /v1/login HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/json

{"username":"user1","password":"user1"}
```
- Ответ
```http request
HTTP/1.1 200 OK
Set-Cookie: ...
Date: ...
```
#### 3. Получить всех пользователей:
```http request
GET /v1/api/user/all HTTP/1.1
Host: 127.0.0.1:8080
Cookie: ...
```
#### 4. Получить пользователя по id:
```http request
GET /v1/api/user/get/{id} HTTP/1.1
Host: 127.0.0.1:8080
Cookie: ...
```
#### 5. Обновить пользователя
```http request
PUT /api/user/update/{id} HTTP/1.1
Host: 127.0.0.1:8080
Cookie: ...
Content-Type: application/json

{"username":"user","email":"user@mail.com","role":1,"address_id":1}
```
#### 5. Удалить пользователя
```http request
DELETE /api/user/delete/{id} HTTP/1.1
Host: 127.0.0.1:8080
Cookie: ...
```


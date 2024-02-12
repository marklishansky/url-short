# url-short (test task)

### grpc & http service for generating short links

### usage

`ghcr.io/marklishansky/url-short:master` - docker image

migrations steps in with startup

#### Переменные окружения

|key|default|
|---|---|
|GRPC_PORT| 10000|
|HTTP_PORT| 10010|
|DB_CONNECTION||

if no `DB_CONNECTION` - inMemory inMemory storage will be used

#### API

`pkg/shorter.swagger.json` - openAPI

`api/shorter.proto` - proto

### deploy

create db migration

```shell
goose -dir sql/migrations create <migration_name> sql
```

when changing schemas as well as updating sql queries you need to run sqlc to generate go from sql.

```shell
make generate-sqlc
```

for code generation from proto

```shell
make generate-grpc-gateway
```
<br />
<br />

### ТЗ

Необходимо реализовать сервис, который должен предоставлять API по созданию сокращённых ссылок следующего формата:

- Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка.
- Ссылка должна быть длинной 10 символов
- Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание)

Сервис должен быть написан на Go и принимать следующие запросы по gRPC:

1. Метод Create, который будет сохранять оригинальный URL в базе и возвращать сокращённый
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL

Решение должно быть предоставлено в «конечном виде», а именно:

- Сервис должен быть распространён в виде Docker-образа
- В качестве хранилища можно использовать in-memory решение или postgresql.
- API должно быть описано в proto файле
- Покрыть реализованный функционал Unit-тестами

### Алгоритм создания ссылки

1. Преобразование ссылки
    1. Хеширование CRC-64
    2. Перевод в 63-ричную ситему счисления в соответствии с заданным алфовитом
    3. Приведение длинны ссылки к 10 символам
2. В случае коллизии к преобразовываемой ссылке прибовляем первый символ алфавита

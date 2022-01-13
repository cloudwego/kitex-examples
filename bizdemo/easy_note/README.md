# Easy Note

---

## Introduction

---
Add a demo for `kitex` which implements a simple note service,the demo is divided into three main sections.

| Service Name    |  Usage    | Framework    | protocol    | Path                   | IDL                                      |
| --------------- | ------------ | ---------- | -------- | ---------------------- | ----------------------------------------- |
| api             | http interface | kitex/gin  | http     | bizdemo/easy_note/api  |                                           |
| kitex.demo.user | user data management | kitex/gorm | protobuf | bizdemo/easy_note/user |        bizdemo/easy_note/idl/user.proto  |
| kitex.demo.note | note data management | kitex/gorm | thrift   | bizdemo/easy_note/note |        bizdemo/easy_note/idl/note.thrift |

### call relations

```
                       ┌───────┐
          ┌───────────►│  api  │◄─────────────────┐
          │            └───────┘                  │
          │                                       │
          │                                       │
          │                                       │
          │                                       │
          │                                       │
┌─────────┴─────────┐                   ┌─────────┴────────┐
│  kitex.demo.user  │                   │ kitex.demo.note  │
└───────────────────┘                   └──────────────────┘
```



###  Use Basic Features

- Middleware、Rate Limiting、Request Retry、Timeout Control、Connection Multiplexing
- Tracing,Customized Access Control
- Service Discovery and Register https://github.com/kitex-contrib/registry-etcd



### catalog introduce

| catalog       | introduce      |
| ---------- | ---------------- |
| constant   | constant        |
| control    |  customized access control    |
| dal,model  | db operation and model     |
| errno      | customized error number |
| middleware | RPC middleware     |
| pack       | data pack         |
| service    | business logic   |



## Quick Start

---
### 1.Setup Basic Dependence
```shell
docker-compose up
```
### 2.Run Note RPC Server
```shell
cd note
sh build.sh
sh output/bootstrap.sh
```

### 3.Run User RPC Server
```shell
cd user
sh build.sh
sh output/bootstrap.sh
```

### 4.Run API Server
```shell
cd api
chmod +x run.sh
./run.sh
```


## API requests 

---
### Register

```shell
curl --location --request POST '127.0.0.1:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"kinggo",
    "password":"123456"
}'
```

### Login
#### will return jwt token

```shell
curl --location --request POST '127.0.0.1:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"kinggo",
    "password":"123456"
}'
```

### Create Note

```shell
curl --location --request POST '127.0.0.1:8080/auth/note' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"test title",
    "content":"test content"
}'
```

### Query Note

```shell
curl --location --request GET '127.0.0.1:8080/auth/note?offset=0&limit=20' \
--header 'Authorization: Bearer $token'
```

### Update Note

```shell
curl --location --request PUT '127.0.0.1:8080/auth/note/$note_id' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"test",
    "content":"test"
}'
```

### Delete Note

```shell
curl --location --request DELETE '127.0.0.1:8080/auth/note/$note_id' \
--header 'Authorization: Bearer $token'
```
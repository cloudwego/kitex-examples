# Easy Note

## Introduction

Add a demo for `kitex` which implements a simple note service,the demo is divided into three main sections.

| Service Name    |  Usage    | Framework    | protocol    | Path                   | IDL                                      |
| --------------- | ------------ | ---------- | -------- | ---------------------- | ----------------------------------------- |
| demoapi         | http interface | kitex/gin  | http     | bizdemo/easy_note/cmd/api  |                                           |
| demouser | user data management | kitex/gorm | protobuf | bizdemo/easy_note/cmd/user |        bizdemo/easy_note/idl/user.proto  |
| demonote | note data management | kitex/gorm | thrift   | bizdemo/easy_note/cmd/note |        bizdemo/easy_note/idl/note.thrift |

### call relations

```
                                    http
                           ┌────────────────────────┐
 ┌─────────────────────────┤                        ├───────────────────────────────┐
 │                         │         demoapi        │                               │
 │      ┌──────────────────►                        │◄──────────────────────┐       │
 │      │                  └───────────▲────────────┘                       │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                           resolve                                 │       │
 │      │                              │                                    │       │
req    resp                            │                                   resp    req
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                   ┌──────────▼─────────┐                          │       │
 │      │                   │                    │                          │       │
 │      │       ┌───────────►       Etcd         ◄─────────────────┐        │       │
 │      │       │           │                    │                 │        │       │
 │      │       │           └────────────────────┘                 │        │       │
 │      │       │                                                  │        │       │
 │      │     register                                           register   │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
┌▼──────┴───────┴───┐                                           ┌──┴────────┴───────▼─┐
│                   │───────────────── req ────────────────────►│                     │
│       demonote    │                                           │        demouser     │
│                   │◄──────────────── resp ────────────────────│                     │
└───────────────────┘                                           └─────────────────────┘
      thrift                                                           protobuf
```

###  Use Basic Features

- Middleware、Rate Limiting、Request Retry、Timeout Control、Connection Multiplexing
- Tracing
  - use jaeger to tracing
- Customized BoundHandler
  - achieve CPU utilization rate customized bound handler
- Service Discovery and Register
  - use [registry-etcd](https://github.com/kitex-contrib/registry-etcd) to discovery and register service

### catalog introduce

| catalog       | introduce      |
| ---------- | ---------------- |
| pkg/constants   | constant        |
| pkg/bound    |  customized bound handler    |
| pkg/errno      | customized error number |
| pkg/middleware | RPC middleware     |
| pkg/tracer  | init jaeger     |
| dal   | db operation              |
| pack       | data pack         |
| service    | business logic   |

## Quick Start

### 1.Setup Basic Dependence
```shell
docker-compose up
```

### 2.Run Note RPC Server
```shell
cd cmd/note
sh build.sh
sh output/bootstrap.sh
```

### 3.Run User RPC Server
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 4.Run API Server
```shell
cd cmd/api
chmod +x run.sh
./run.sh
```

### 5.Jaeger 

visit `http://127.0.0.1:16686/` on  browser.

#### Snapshots

<img src="images/shot.png" width="2850"  alt=""/>

## API requests 

### Register
```shell
curl --location --request POST '127.0.0.1:8080/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"kinggo",
    "password":"123456"
}'
```

#### response
```javascript
{
    "code": 0,
    "message": "Success",
    "data": null
}
```

### Login

#### will return jwt token
```shell
curl --location --request POST '127.0.0.1:8080/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"kinggo",
    "password":"123456"
}'
```

#### response
```javascript
{
    "code": 200,
    "expire": "2022-01-19T01:56:46+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI1Mjg2MDYsImlkIjoxLCJvcmlnX2lhdCI6MTY0MjUyNTAwNn0.k7Ah9G4Enap9YiDP_rKr5HSzF-fc3cIxwMZAGeOySqU"
}
```

### Create Note
```shell
curl --location --request POST '127.0.0.1:8080/v1/note' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"test title",
    "content":"test content"
}'
```

#### response
```javascript
{
    "code": 0,
    "message": "Success",
    "data": null
}
```

### Query Note
```shell
curl --location --request GET '127.0.0.1:8080/v1/note/query?offset=0&limit=20&search_keyword=test' \
--header 'Authorization: Bearer $token'
```

#### response
```javascript
{
    "code": 0,
    "message": "Success",
    "data": {
        "notes": [
            {
                "note_id": 1,
                "user_id": 1,
                "user_name": "kinggo",
                "user_avatar": "test",
                "title": "test title",
                "content": "test content",
                "create_time": 1642525063
            }
        ],
        "total": 1
    }
}
```

### Update Note
```shell
curl --location --request PUT '127.0.0.1:8080/v1/note/$note_id' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"test",
    "content":"test"
}'
```

#### response
```javascript
{
    "code": 0,
    "message": "Success",
    "data": null
}
```

### Delete Note
```shell
curl --location --request DELETE '127.0.0.1:8080/v1/note/$note_id' \
--header 'Authorization: Bearer $token'
```

#### response
```javascript
{
    "code": 0,
    "message": "Success",
    "data": null
}
```

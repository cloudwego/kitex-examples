# Easy Note

## Introduction

Add a demo for `kitex` which implements a simple note service,the demo is divided into three main sections.

| Service Name    |  Usage    | Framework    | protocol    | Path                   | IDL                                      |
| --------------- | ------------ | ---------- | -------- | ---------------------- | ----------------------------------------- |
| demoapi         | http interface | kitex/hertz  | http     | bizdemo/easy_note/cmd/api  |                                           |
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

## Custom Error Code

Customise the response error code in the `errno` package.

```go
const (
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
)
```

Sample code : Replace the default error code for hertz-jwt authentication error with a custom error code.

```go
authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
    Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
        c.JSON(code, map[string]interface{}{
            "code":    errno.AuthorizationFailedErrCode,
            "message": message,
        })
    },
    //Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
    //  c.JSON(code, map[string]interface{}{
    //      "code":    code,
    //      "message": message,
    //  })
    //}
})
```

## API requests

[API requests](api.md)


## Faq

### How to upgrade kitex_gen

- refer to [Makefile](Makefile)
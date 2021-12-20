#Easy Note

---
## Quick Start

---
### 1.Setup Basic Dependence
```shell
docker-compose up
```
### 2.Run Note RPC Server
```shell
sh note/build.sh
sh note/output/bootstrap.sh
```

### 3.Run User RPC Server
```shell
sh user/build.sh
sh user/output/bootstrap.sh
```

### 4.Run API Server
```shell
sh api/run.sh
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
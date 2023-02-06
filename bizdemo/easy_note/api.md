## API requests

The following is a list of API requests and partial responses.

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
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}
// failed
{
    "code": 10003,
    "message": "User already exists",
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
// successful
{
    "code": 0,
    "expire": "2022-01-19T01:56:46+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI1Mjg2MDYsImlkIjoxLCJvcmlnX2lhdCI6MTY0MjUyNTAwNn0.k7Ah9G4Enap9YiDP_rKr5HSzF-fc3cIxwMZAGeOySqU"
}
// failed
{
    "code": 10004,
    "message": "Authorization failed",
    "data": null
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
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}
// failed
{
    "code": 10002,
    "message": "Wrong Parameter has been given",
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
// successul
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
// failed
{
    "code":10002,
    "message":"Wrong Parameter has been given",
    "data":null
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
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}
// failed
{
    "code":10001,
    "message":"strconv.ParseInt: parsing \"$note_id\": invalid syntax",
    "data":null
}
```

### Delete Note
```shell
curl --location --request DELETE '127.0.0.1:8080/v1/note/$note_id' \
--header 'Authorization: Bearer $token'
```

#### response
```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}
// failed
{
    "code":10001,
    "message":"strconv.ParseInt: parsing \"$note_id\": invalid syntax",
    "data":null
}
```

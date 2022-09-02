# Forward Meta Information Transmitting

## Running the example

### running server-2

`server-2` will be called by `server-1`, so we must run it firstly.
`server-2` checks the metainfo from `server-1`.

```
go run server-2/main.go
```

### running server-1

`server-1` will be called by `client`.
`server-1` checks the metainfo from `client`.

```
go run server-1/main.go
```

### running client

run `client`,and set some metainfo into context.

```
go run client/main.go
```

## NOTE

- Must use certain transport protocols that can transmit meta information,such as TTHeader, HTTP.
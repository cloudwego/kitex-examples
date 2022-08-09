# sentinel/kitex

This is an example of using in `sentinel-golang` in `kitex`.

The extensions are implemented as middleware,
but Kitex provides additional interfaces to support fusing and flow limiting, and Sentinel's extensions do not 
necessarily take effect when both are used at the same time.

server:
- Example of using sentinel-golang in kitex server.
- You can do this by executing `go run ./server/main.go` command in the terminal to see the results.

client:
- Example of using sentinel-golang in kitex client.
- You can do this by executing `go run ./client/main.go` command in the terminal to see the results.


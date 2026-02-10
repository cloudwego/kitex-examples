# kitex_zorm

## Introduce

A demo with `Kitex` and `Zorm`

- Use `thrift` IDL to define `RPC` interface
- Use `kitex` to generate code
- Use `Zorm` and `MySQL`

## IDL

This demo use `thrift` IDL to define `RPC` interface. The specific interface define in [user.thrift](idl/user.thrift)

## Code generation tool

This demo use `kitex` to generate code. The use of `kitex` refers
to [kitex](https://www.cloudwego.io/docs/kitex/tutorials/code-gen/)

The `kitex` commands used can be found in [Makefile](Makefile)

## Zorm

This demo use `Zorm` to operate `MySQL` and refers to [Zorm](https://www.zorm.cn/)

## How to run

### Run mysql docker

```bash
cd bizdemo/kitex_zorm && docker compose up
```

### Run demo

```bash
cd bizdemo/kitex_zorm
go build -o kitex_zorm && ./kitex_zorm
```

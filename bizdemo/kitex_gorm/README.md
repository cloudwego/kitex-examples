# kitex_gorm

## Introduce

A demo with `Kitex` and `Gorm`

- Use `thrift` IDL to define `RPC` interface
- Use `kitex` to generate code
- Use `Gorm` and `MySQL`

## IDL

This demo use `thrift` IDL to define `RPC` interface. The specific interface define in [user.thrift](idl/user.thrift)

## Code generation tool

This demo use `kitex` to generate code. The use of `kitex` refers
to [kitex](https://www.cloudwego.io/docs/kitex/tutorials/code-gen/)

The `kitex` commands used can be found in [Makefile](Makefile)

## Gorm

This demo use `Gorm` to operate `MySQL` and refers to [Gorm](https://gorm.io/)

## How to run

### Run mysql docker

```bash
cd bizdemo/kitex_gorm && docker compose up
```

### Run demo

```bash
cd bizdemo/kitex_gorm
go build -o kitex_gorm && ./kitex_gorm
```

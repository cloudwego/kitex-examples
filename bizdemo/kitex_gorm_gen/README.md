# kitex_gorm

## Introduce

A demo with `Kitex` and `Gorm`

- Use `thrift` IDL to define `RPC` interface
- Use `kitex` to generate code
- Use `Gorm/Gen` and `MySQL`

- `/cmd` generate gorm_gen
- `/dao` initialize mysql connection
- `/model` gorm generated user struct
- `/hander.go` basic biz logic for updateUser, addUser, deleteUser, queryUser

## IDL

This demo use `thrift` IDL to define `RPC` interface. The specific interface define in [user.thrift](idl/user.thrift)

## Code generation tool

This demo use `kitex` to generate code. The use of `kitex` refers
to [kitex](https://www.cloudwego.io/docs/kitex/tutorials/code-gen/)

The `kitex` commands used can be found in [Makefile](Makefile)

## GORM/Gen

GEN: Friendly & Safer GORM powered by Code Generation.

This demo use `GORM/Gen` to operate `MySQL` and refers to [Gen](https://gorm.io/gen/index.html).

#### Quick Start

- Update the Database DSN to your own in [Database init file](dao/mysql/init.go).
- Refer to the code comments, write the configuration in [Generate file](cmd/generate.go).
- Using the following command for code generation, you can generate structs from databases or basic type-safe DAO API for struct.

```bash
cd bizdemo/kitex_gorm_gen/cmd
go run generate.go
```

- For more Gen usage, please refer to [Gen Guides](https://gorm.io/gen/index.html).

## How to run

### Run mysql docker

```bash
cd bizdemo/kitex_gorm_gen && docker compose up
```

### Run demo

```bash
cd bizdemo/kitex_gorm_gen
go build -o kitex_gorm_gen && ./kitex_gorm_gen
```

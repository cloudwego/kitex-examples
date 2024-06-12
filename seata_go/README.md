# Kitex with Seata-go example

This example uses the [use case](https://seata.apache.org/docs/user/quickstart/#use-case) from the Seata official documentation to demonstrate how to use Seata-go in Kitex.

## How To Run

### Basic environment

Ensure Docker has been installed.

```shell
cd dockercompose
docker-compose up -d
```

### Run all servers

```shell
go run ./service/account
go run ./service/storage
go run ./service/order
```

### Run main program

```shell
go run main.go
```

## Basic Principle

By leveraging Kitex's capability to transmit [metainfo](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/metainfo/) between services, the XID assigned by Seata to each transaction can be propagated to each service. This XID can then be used to link the entire transaction.

In this example, Kitex middleware is used to automatically propagate the XID. Please refer to [seata-go.go](./middleware/seata-go.go).

**NODE**: Must use the underlying transport protocol that supports passthrough of meta information

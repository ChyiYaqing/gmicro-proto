# gmicro-proto

创建`golang/order/v0.0.1`标签tag

```bash
./run.sh order v0.0.1 chyiyaqing chyiyaqing@gmail.com
```

创建`golang/payment/v0.0.1`标签tag

```bash
./run.sh payment v0.0.1 chyiyaqing chyiyaqing@gmail.com
```

创建`golang/shipping/v0.0.1`标签tag

```bash
./run.sh shipping v0.0.1 chyiyaqing chyiyaqing@gmail.com
```

打Tag

```bash
git tag v0.0.8 && git push origin v0.0.8
```

## Requisition

| Tool | Description | Installation |
|:-----|:------------|:-------------|
| ptotobuf | protocol buffer | [install protoc](https://google.github.io/proto-lens/installing-protoc.html) |
| protoc-gen-go | Plugin for the Google protocol buffer compiler to generate Go code. | [install protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/) |
| protoc-gen-go-grpc | This project aims to provide that HTTP+JSON interface to your gRPC service. | [install protoc-gen-go-grpc](https://grpc.io/docs/languages/go/quickstart/) |
| protoc-gen-grpc-gateway | Plugin for Google protocol buffer compiler to generate a reverse-proxy, which converts incoming RESTful HTTP/1 requests gRPC invocation. | [install grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway#installation) |
| protoc-gen-openapiv2 | Plugin for Google protocol buffer compiler to generate open API config file. | [install openapiv2](https://github.com/grpc-ecosystem/grpc-gateway#installation) |
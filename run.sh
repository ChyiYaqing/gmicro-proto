#!/bin/bash

SERVICE_NAME=$1
RELEASE_VERSION=$2
USER_NAME=$3
EMAIL=$4

git config user.name "$USER_NAME"
git config user.email "$EMAIL"
git fetch --all && git checkout main

# 安装依赖
sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest 
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest 
go install github.com/bufbuild/buf/cmd/buf@latest

echo "Checking installed binaries:"
ls -l $(go env GOPATH)/bin/
which protoc-gen-go
which protoc-gen-grpc-gateway

# 生成protobuf代码,将.proto文件编译成Go代码，包括gRPC服务代码
# 每一个 -I 标志代表一个用于搜索导入的目录
protoc -I ./third_party -I ./proto \
    --go_out=./golang \
    --go_opt=paths=source_relative \
    --go-grpc_out=./golang \
    --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./golang \
    --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt=generate_unbound_methods=true \
    --openapiv2_out=./openapiv2/OpenAPI \
    --openapiv2_opt=generate_unbound_methods=true \
    ${SERVICE_NAME}/**/*.proto
#
# 选用buf 执行
#buf dep update
#buf generate

# 初始化Go 模块
mkdir -p golang/${SERVICE_NAME}
cd golang/${SERVICE_NAME}
go mod init github.com/chyiyaqing/gmicro-proto/golang/${SERVICE_NAME} || true
go mod tidy

cd ../../
git pull && git add . && git commit -am "proto update" || true
git push origin HEAD:main

# 版本标签管理
# -f: force 强制创建标签，如果标签存在则覆盖
# -a: annotated 创建一个带注释的标签
# -m: 指定标签的描述信息
git tag -fa golang/${SERVICE_NAME}/${RELEASE_VERSION} -m "golang/${SERVICE_NAME}/${RELEASE_VERSION}"
git pull && git push origin refs/tags/golang/${SERVICE_NAME}/${RELEASE_VERSION}
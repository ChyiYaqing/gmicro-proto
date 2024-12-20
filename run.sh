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

# 生成protobuf代码,将.proto文件编译成Go代码，包括gRPC服务代码
protoc \
    --go_out=./golang \
    --go_opt=paths=source_relative \
    --go-grpc_out=./golang \
    --go-grpc_opt=paths=source_relative \
    ./${SERVICE_NAME}/*.proto

# 初始化Go 模块
cd golang/${SERVICE_NAME}
go mod init github.com/chyiyaqing/gmicro-proto/golang/${SERVICE_NAME} || true
go mod tidy

cd ../../
git add . && git commit -am "proto update" || true
git push origin HEAD:main

# 版本标签管理
# -f: force 强制创建标签，如果标签存在则覆盖
# -a: annotated 创建一个带注释的标签
# -m: 指定标签的描述信息
git tag -fa golang/${SERVICE_NAME}/${RELEASE_VERSION} -m "golang/${SERVICE_NAME}/${RELEASE_VERSION}"
git push origin refs/tags/golang/${SERVICE_NAME}/${RELEASE_VERSION}
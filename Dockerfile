# syntax = docker/dockerfile:experimental

# 构建阶段
FROM golang:1.24.7-alpine AS builder

# 设置环境变量
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置工作目录
WORKDIR /app

# 构建依赖
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod verify

# 程序编译
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build \
    go build -v -trimpath -ldflags "-s -w" -o ./main ./main.go

# 运行阶段
FROM alpine:3.21

# 安装证书和时区数据
RUN apk add --no-cache ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

# 设置工作目录
WORKDIR /app

# 拷贝配置
COPY --from=builder /app/conf ./conf

# 拷贝应用程序
COPY --from=builder /app/main ./main

# 暴露端口
EXPOSE 8888

# 启动应用程序
ENTRYPOINT ["./main"]
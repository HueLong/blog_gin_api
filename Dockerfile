# 构建阶段
FROM golang:1.20 AS builder

# 设置工作目录
WORKDIR /app

# 设置 GOPROXY 使用国内代理
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# 首先只复制依赖文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用（使用单核编译，减少 CPU 使用率）
RUN go build -ldflags="-s -w" -trimpath -o blog_gin_api ./cmd/api

# 运行阶段
FROM debian:bullseye-slim

# 设置时区
ENV TZ=Asia/Shanghai
RUN apt-get update && apt-get install -y ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/blog_gin_api .
COPY --from=builder /app/config ./config

# 创建日志目录
RUN mkdir -p /app/logs

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./blog_gin_api"] 
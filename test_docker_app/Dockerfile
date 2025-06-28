# 阶段1：构建环境
FROM golang:1.23.5-alpine AS builder

# 设置构建参数
ARG APP_VERSION=v1.0.0

# 设置工作目录
WORKDIR /app

# 复制依赖文件
COPY go.mod .
# COPY go.sum .

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 \ 
    # 禁用CGO,禁止使用C语言扩展，纯Go编译
    GOOS=linux \ 
    # GO Operator System，强制生成linux可执行文件
    go build -ldflags="-X 'main.Version=${APP_VERSION}'" -o main .
    # -ldflags 高级链接参数，-X:允许在编译时注入变量值
    # main.Version 表示要注入的变量路径（包名.变量名）
    # ${APP_VERSION} 来自 Docker 的 ARG 或 --build-arg

# 安装CA证书（HTTPS支持）
# RUN apk --no-cache add ca-certificates
# 先不用，比较大

# 暴露端口
EXPOSE 8080

# 设置环境变量
ENV APP_VERSION=1.0.0

# 启动命令
# CMD ["./main"]
CMD ["go", "run", "."]
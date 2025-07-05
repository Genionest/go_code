# 构建阶段
# FROM golang:1.21-alpine AS builder
FROM golang:1.23.5-alpine
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct \
    GOPATH=/go
WORKDIR /app
COPY . .
RUN go mod download
# RUN go build -o main .
CMD ["go", "run", "."]

# 运行阶段
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /app/main .
# CMD ["./main"]
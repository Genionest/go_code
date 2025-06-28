# 构建镜像（带版本参数）
docker build --build-arg APP_VERSION=v1.2.0 -t go-app:latest .
# -t 设置标签 latest就是设置的标签

# 运行容器
docker run -d -p 8080:8080 --name go-container -e APP_VERSION=1.2.0 go-app:latest
# -d 后台运行
# -p 端口映射 本机端口:容器端口
# --name 容器名，不指定就只有镜像名，容器名是docker给的，如果没有，docker会随机命名

# 测试访问
curl http://localhost:8080
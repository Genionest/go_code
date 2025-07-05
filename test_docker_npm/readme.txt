# 构建
docker build -t my-npm-app .

# 创建容器并运行
docker run -d -p 3001:3000 --name my-running-npm my-npm-app

# 验证
curl http://localhost:3001
# 输出：Hello from Dockerized npm!

# 查看日志
docker logs my-running-npm
# 输出：Server running on http://localhost:3000
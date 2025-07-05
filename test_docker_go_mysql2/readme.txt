# 创建专用网络
docker network create app-network

# 构建并启动mysql容器
docker run -d --name mysql-container \
  --network app-network \
  -e MYSQL_ROOT_PASSWORD=mysecret \
  -e MYSQL_DATABASE=mydb \
  -v mysql-data:/var/lib/mysql \
  mysql:8.0
# -e xxx  设置环境变量
# --network xxx  连接到xxx的docker网络

# 构建golang镜像
docker build -t golang-app .

# 运行golang容器
docker run -it \
  --name golang-container \
  --network app-network \
  golang-app
# -i  保持标准输入流打开，允许与容器交互
# -t  分配一个伪终端，使容器像本地终端一样工作
# --rm  容器停止后直接删除容器, 这里我们没有使用
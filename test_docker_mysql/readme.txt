# 1. 构建自定义镜像
docker build -t custom-mysql .

# 2. 运行容器（带数据持久化）
docker run -d \
  --name mysql-container \
  -p 3307:3306 \
  -v mysql_data:/var/lib/mysql \
  custom-mysql

# 3. 验证运行
docker ps --filter name=mysql-container


通过容器内命令行连接
docker exec -it mysql-container mysql -u app_user -papp_password
# -p app_password 后面的是密码

# 外部客户端连接
mysql -h 127.0.0.1 -u app_user -papp_password

-- 验证初始化结果
SHOW DATABASES;
USE app_db;
DESCRIBE users;
SELECT user FROM mysql.user;
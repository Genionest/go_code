# 使用官方 MySQL 8.0 镜像作为基础
FROM mysql:8.0

# 设置环境变量
ENV MYSQL_ROOT_PASSWORD=my_root_password
ENV MYSQL_DATABASE=app_db
ENV MYSQL_USER=app_user
ENV MYSQL_PASSWORD=app_password

# 复制配置文件到容器
COPY my.cnf /etc/mysql/conf.d/custom.cnf

# 复制初始化脚本
COPY init.sql /docker-entrypoint-initdb.d/

# 设置数据存储卷
VOLUME /var/lib/mysql

# 暴露 MySQL 端口
EXPOSE 3306

# 设置容器启动命令
CMD ["mysqld"]
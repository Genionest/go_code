# 使用官方 PostgreSQL 镜像
FROM postgres:15-alpine

# 设置环境变量（可选）
ENV POSTGRES_USER=myuser \
    POSTGRES_PASSWORD=mypassword \
    POSTGRES_DB=mydatabase
# 上面的内容，可以随意命名，不固定
# 设置完这些环境变量后，会自动创建对应的用户，密码，数据库，

# 将初始化脚本复制到容器中
COPY ./docker-entrypoint-initdb.d /docker-entrypoint-initdb.d

# 暴露 PostgreSQL 默认端口
EXPOSE 5432
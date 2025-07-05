# 构建镜像
docker build -t my-postgres .

# 创建容器并运行
docker run -d --name my-postgres-container \
  -p 5433:5432 \
  -v pgdata:/var/lib/postgresql/data \
  my-postgres


# 验证数据
# 进入容器中的 PostgreSQL
docker exec -it my-postgres-container psql -U myuser -d mydatabase

# 执行 SQL 查询
SELECT * FROM users;

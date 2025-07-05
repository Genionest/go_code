# 使用官方 Node.js 镜像（Alpine 轻量版）
FROM node:20-alpine

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY package*.json ./

# 安装依赖（生产环境依赖）
# RUN npm ci --only=production # 只能在有package-lock.json时使用RUN npm ci
RUN npm install

# 复制应用源代码
COPY . .

# 暴露端口
EXPOSE 3000

# 启动命令
CMD ["npm", "start"]
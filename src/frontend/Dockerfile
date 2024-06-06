# 基于 Node 的官方镜像
FROM node:lts-alpine

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY package*.json ./

# 安装项目依赖
RUN npm install

# 复制项目文件
COPY . .

# 构建应用
RUN npm run build

# 安装 serve 来运行构建的应用
RUN npm install -g serve

# 启动服务
CMD ["serve", "-s", "dist"]

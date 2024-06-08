# 启动前端

cd src/frontend
docker build -t frontend .
docker run -p 8080:3000 frontend
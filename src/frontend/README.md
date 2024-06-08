# frontend

## 项目环境配置

### 直接运行

执行如下命令配置前端运行环境：

```
npm install
```

`npm` 会自动完成依赖配置，之后使用命令：

```
npm run serve
```

运行前端服务，并在浏览器中查看。

### Docker 运行

使用 Docker 运行前端界面，切换到`\frontend`文件夹下，执行如下指令：

```
docker build -t frontend .
docker run -p 8080:3000 frontend
```

之后在本机的 8080 端口即可完成访问。

### 运行 IP 配置

对于后端不同的服务器 IP 地址，可以在`src/.env`中进行配置。后端 API 访问的 IP 地址需修改一下字段：

```
VUE_APP_API_URL = http://localhost:8081
```

重新启动并运行。

### Compiles and hot-reloads for development

```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

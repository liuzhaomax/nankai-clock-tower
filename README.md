# 南开钟楼

1. 打包
```shell
make build
```
2. 新建nct文件夹
3. 拷贝bin, environment, Dockerfile, docker-compose, .env 
4. cd到nct，构建镜像
```shell
# 如果直接构建失败，先尝试手动拉取镜像
docker pull alpine:3.18 --platform linux/amd64
# 构建镜像
docker build -t nct .
```
5. 创建容器
```shell
docker compose up -d
```

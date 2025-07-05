# 南开钟楼

1. 新建nct文件夹
2. 拷贝bin，environment，Dockerfile，docker-compose
3. cd到nct，构建镜像
```shell
docker build -t nct .
```
4. 创建容器
```shell
docker compose up -d
```

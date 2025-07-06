# 南开钟楼

1. 打包
```shell
make build
```
2. 新建nct文件夹
3. 拷贝bin，environment，Dockerfile，docker-compose 
4. cd到nct，构建镜像
```shell
docker build -t nct .
```
5. 创建容器
```shell
docker compose up -d
```

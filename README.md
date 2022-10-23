# video_server

## video_server介绍
视频后台

## video_server项目功能点
- [x] 单点登陆
- [x] 全局搜索
- [x] 用户管理
- [x] 角色权限
- [x] 视频上传
- [x] 分类管理
- [x] 视频信息管理
- [x] 定时任务删除视频

## 快速启动
将`config.yaml`中`addr`改为mysqldb,否则无法与mysql容器连接
将服务地址改为0.0.0.0，否则无法从容器外访问
```bash
docker run -d --rm --name mysqldb -p 3306:3306 hy6w/mariadb:latest
docker build -t video_server:v1.0.0 .
docker run -d --rm -p 8080:8080 -v E:\Desktop\video_server:/app/video_server --name video_server_project --link MySQL:mysqldb video_server:v1.0.0
```

## 生成Swagger
```shell
bash ./swagger.sh
```
http://localhost:8080/swagger/index.html



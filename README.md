# video_server
视频后台

## 实现功能
- [x] redis实现单点登陆
- [x] cobra启动
- [x] cabin权限认证
- [x] 定时任务删除视频文件，每天晚上3点删除视频
- [x] 完善配置文件
- [x] swagger

## 待完成功能
- [x] 日志写入log文件

## 快速启动
将config.yaml中MysqlHost改为mysqldb,否则无法与mysql容器连接
将服务地址改为0.0.0.0，否则无法从容器外访问
```bash
docker build -t video_server:v1.0.0 .
```

```bash
docker run -d --rm -p 8080:8080 -v E:\Desktop\video_server:/app/video_server --name video_server_project --link MySQL:mysqldb video_server:v1.0.0
```

## 访问 swagger文档
[video_server Swagger](http://localhost:8080/swagger/index.html#/)

### 生成swagger时注意事项
#### 1.如果api不在main.go中，要指定api，-g指定要扫描的文件(基础信息文件和定义的api文件)，-o指定输出文件夹
```bash
swag init -g cmd/main.go router/endpoint.go -o docs
```
#### 2.api文件中要导入docs目录

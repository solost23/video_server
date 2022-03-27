# video_server
视频后台

## 实现功能
- [x] cabin权限认证
- [x] 定时任务删除视频文件，每天晚上3点删除视频

## 待完成功能

- [x] swagger
- [x] 完善配置文件
- [x] 日志写入log文件

### 生成swagger时注意事项
#### 1.如果api不在main.go中，要指定api，-g指定要扫描的文件(基础信息文件和定义的api文件)，-o指定输出文件夹
```bash
swag init -g cmd/main.go router/endpoint.go -o docs
```
#### 2.api文件中要导入docs目录
package initialize

import (
	"github.com/solost23/tools/minio_storage"
	"video_server/global"
)

func InitMinio() {
	var err error
	global.Minio, err = minio_storage.NewMinio(&minio_storage.Config{
		EndPoint:        global.ServerConfig.MinioConfig.EndPoint,
		AccessKeyID:     global.ServerConfig.MinioConfig.AccessKeyId,
		SecretAccessKey: global.ServerConfig.MinioConfig.SecretAccesskey,
		UserSSL:         global.ServerConfig.MinioConfig.UserSSL,
	})
	if err != nil {
		panic(err)
	}
}

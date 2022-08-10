package workList

import (
	"context"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"path"
	"strconv"
	"time"
	"video_server/config"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"github.com/minio/minio-go"

	"github.com/solost23/tools/minio_storage"
)

func UploadImg(user *models.User, folderName string, file *multipart.FileHeader) (string, error) {
	return uploadImage(user, folderName, file, "image")
}

func UploadVid(user *models.User, folderName string, file *multipart.FileHeader) (string, error) {
	return uploadImage(user, folderName, file, "video")
}

func uploadImage(user *models.User, folderName string, file *multipart.FileHeader, uploadType string) (string, error) {
	fileHandle, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func() { _ = fileHandle.Close() }()

	fileByte, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		return "", err
	}

	postFileName := file.Filename
	fileName := utils.GetMd5Hash(
		time.Now().Format(models.TimeFormat)+
			strconv.Itoa(int(user.ID))+
			utils.GetMd5Hash(string(fileByte))+
			postFileName) + path.Ext(postFileName)
	url, err := upload(folderName, fileName, fileHandle, uploadType)
	if err != nil {
		return "", err
	}
	return url, nil
}

func upload(folderName string, fileName string, fileHandle multipart.File, uploadType string) (result string, err error) {
	minioConfig := config.NewMinio()
	client, err := minio_storage.NewMinio(&minio_storage.Config{
		EndPoint:        minioConfig.EndPoint,
		AccessKeyID:     minioConfig.AccessKeyID,
		SecretAccessKey: minioConfig.SecretAccessKey,
		UserSSL:         minioConfig.UserSSL,
	})
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	if err = minio_storage.CreateBucket(ctx, client, folderName); err != nil {
		return "", err
	}
	//err = minio_storage.FileUpload(ctx, client, "bucket1", fileName, folderName, "Application/"+uploadType)
	//if err != nil {
	//	return "", err
	//}
	_, err = client.PutObjectWithContext(ctx, folderName, fileName, fileHandle, 100, minio.PutObjectOptions{ContentType: "Application/" + uploadType})
	if err != nil {
		return "", err
	}
	// 上传文件流数据
	requestParams := make(url.Values)
	fileUrl, err := minio_storage.GetFileUrl(ctx, client, folderName, fileName, 300*time.Hour, requestParams)
	if err != nil {
		return "", err
	}
	return fileUrl, nil
}

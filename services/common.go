package services

import (
	"context"
	"github.com/solost23/go_interface/gen_go/common"
	"github.com/solost23/go_interface/gen_go/oss"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strconv"
	"time"
	"video_server/global"
	"video_server/pkg/constants"
	"video_server/pkg/models"
	"video_server/pkg/utils"
)

func UploadImg(ctx context.Context, user *models.User, folderName string, file *multipart.FileHeader) (string, error) {
	return uploadImage(ctx, user, folderName, file, "image")
}

func UploadVid(ctx context.Context, user *models.User, folderName string, file *multipart.FileHeader) (string, error) {
	return uploadImage(ctx, user, folderName, file, "video")
}

func uploadImage(ctx context.Context, user *models.User, folderName string, file *multipart.FileHeader, uploadType string) (string, error) {
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
		time.Now().Format(constants.TimeFormat)+
			strconv.Itoa(int(user.ID))+
			utils.GetMd5Hash(string(fileByte))+
			postFileName) + path.Ext(postFileName)
	reply, err := global.OSSSrvClient.Upload(ctx, &oss.UploadRequest{
		Header: &common.RequestHeader{
			OperatorUid: int32(user.ID),
			Requester:   user.UserName,
		},
		Folder:     folderName,
		Key:        fileName,
		Data:       fileByte,
		UploadType: uploadType,
	})
	if err != nil {
		return "", err
	}
	return reply.GetUrl(), nil
}

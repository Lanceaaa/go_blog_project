package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

// 上传文件工具库
func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	// 获取文件名
	fileName := upload.GetFileName(fileHeader.Filename)
	// 获取上传文件的最终保存目录
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	// 判断文件后缀
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	// 判断文件保存目录是否存在
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	// 判断文件大小是否超出
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}
	// 判断文件是否拥有权限
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}
	// 保存文件
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	// 返回保存的路径
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}

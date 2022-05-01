package shrd_service

import (
	"context"
	"mime/multipart"
)

type FileSvc interface {
	UploadPrivateFile(ctx context.Context, file multipart.File, path string) (string, error)
	UploadPublicFile(ctx context.Context, file multipart.File, path string) (string, error)
	DeleteFile(ctx context.Context, bucketName string, path string) error
	GetPreSignUrl(ctx context.Context, path string) (string, error)
	BuildPublicUrl(path string) string
}

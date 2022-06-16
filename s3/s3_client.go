package s3_client

import (
	"context"
	"fmt"
	"mime/multipart"

	shrd_service "github.com/StevanoZ/dv-shared/service"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Init(baseConfig *shrd_utils.BaseConfig) (*s3.Client, error) {
	creeds := credentials.NewStaticCredentialsProvider(baseConfig.AWSAccessKey, baseConfig.AWSSecretKey, "")
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(creeds),
		config.WithRegion(baseConfig.AWSRegion),
	)

	return s3.NewFromConfig(cfg), err
}

func PreSignClient(client *s3.Client) *s3.PresignClient {
	return s3.NewPresignClient(client)
}

type S3ClientImpl struct {
	client  S3Client
	preSign S3PreSign
	config  *shrd_utils.BaseConfig
}

func NewS3Client(
	client S3Client,
	preSign S3PreSign,
	config *shrd_utils.BaseConfig,
) shrd_service.FileSvc {
	return &S3ClientImpl{
		client:  client,
		preSign: preSign,
		config:  config,
	}
}

func (s *S3ClientImpl) UploadPrivateFile(ctx context.Context, file multipart.File, path string) (string, error) {
	_, err := s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.config.S3PrivateBucketName),
		Key:    aws.String(path),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	shrd_utils.LogInfo(fmt.Sprintf("success uploaded file to S3: %s", path))

	preSignUrl, err := s.GetPreSignUrl(ctx, path)
	if err != nil {
		return "", err
	}

	return preSignUrl, nil
}

func (s *S3ClientImpl) UploadPublicFile(ctx context.Context, file multipart.File, path string) (string, error) {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.config.S3PublicBucketName),
		Key:         aws.String(path),
		Body:        file,
		ContentType: aws.String(shrd_utils.GetExt(path)),
	})
	if err != nil {
		return "", err
	}

	return s.BuildPublicUrl(path), nil
}

func (s *S3ClientImpl) DeleteFile(ctx context.Context, bucketName string, path string) error {
	_, err := s.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	})

	return err
}

func (s *S3ClientImpl) GetPreSignUrl(ctx context.Context, path string) (string, error) {
	params := &s3.GetObjectInput{
		Bucket:              aws.String(s.config.S3PrivateBucketName),
		Key:                 aws.String(path),
		ResponseContentType: aws.String(shrd_utils.GetExt(path)),
	}

	resp, err := s.preSign.PresignGetObject(ctx, params, func(po *s3.PresignOptions) {
		po.Expires = s.config.PreSignUrlDuration
	})
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}

func (s *S3ClientImpl) BuildPublicUrl(path string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s",
		s.config.S3PublicBucketName,
		s.config.AWSRegion,
		path,
	)
}

package s3client

import (
	"context"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	ListObjectsV2(ctx context.Context,
		params *s3.ListObjectsV2Input,
		optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	DeleteObject(ctx context.Context,
		params *s3.DeleteObjectInput,
		optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)

	// UNCOMMENT IF NEEDED
	// ListBuckets(ctx context.Context,
	// 	params *s3.ListBucketsInput,
	// 	optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	// CopyObject(ctx context.Context,
	// 	params *s3.CopyObjectInput,
	// 	optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error)
	// CreateBucket(ctx context.Context,
	// 	params *s3.CreateBucketInput,
	// 	optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
	// DeleteBucket(ctx context.Context,
	// 	params *s3.DeleteBucketInput,
	// 	optFns ...func(*s3.Options)) (*s3.DeleteBucketOutput, error)
	// GetBucketAcl(ctx context.Context,
	// 	params *s3.GetBucketAclInput,
	// 	optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
	// GetObjectAcl(ctx context.Context,
	// 	params *s3.GetObjectAclInput,
	// 	optFns ...func(*s3.Options)) (*s3.GetObjectAclOutput, error)
}

type S3PreSign interface {
	PresignGetObject(
		ctx context.Context,
		params *s3.GetObjectInput,
		optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

type ListObjectsV2Pager interface {
	HasMorePages() bool
	NextPage(context.Context, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}
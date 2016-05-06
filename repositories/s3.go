package repositories

import (
	"github.com/bunniesandbeatings/pivnet-cli/rc"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
)

type S3 struct {
	session *s3.S3
	bucket string
}

func NewS3Repository(repositoryDefinition rc.Repository) *S3 {
	session := s3.New(
		session.New(
			&aws.Config{
				Region: aws.String(repositoryDefinition.Region),
				Credentials: credentials.NewStaticCredentials(
					repositoryDefinition.AccessKeyId,
					repositoryDefinition.SecretAccessKey,
					"",
				),
			}))

	repository := S3{
		session: session,
		bucket: repositoryDefinition.Bucket,
	}

	return &repository
}

func (repository *S3) Search() error {
	prefix := "product_files/"

	response, err := repository.session.ListObjects(
		&s3.ListObjectsInput{
			Bucket: aws.String(repository.bucket),
			Prefix: &prefix,
		})

	if err != nil {
		fmt.Errorf("could not list objects at prefix %s", prefix)
		return err
	}

	for _, key := range response.Contents {
		fmt.Println(*key.Key)
	}

	return nil
}
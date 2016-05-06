package commands

import (
	"github.com/bunniesandbeatings/pivnet-cli/rc"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type SearchCommand struct {

}

func (command *SearchCommand) Execute(args []string) error {
	config, err := rc.Read()
	if err != nil {
		return err
	}

	fmt.Println("Listing objects...\n")

	service := s3.New(
		session.New(
			&aws.Config{
				Region: aws.String(config.S3.Region),
				Credentials: credentials.NewStaticCredentials(
					config.S3.AccessKeyId, config.S3.SecretAccessKey,
					"",
				),
			}))

	prefix := "product_files/"

	response, err := service.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(config.S3.Bucket),
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

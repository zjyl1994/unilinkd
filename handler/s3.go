package handler

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/zjyl1994/unilinkd/config"
)

func S3Handler(c *fiber.Ctx, url string) error {
	if config.S3 == nil {
		return errors.New("s3 not configure")
	}
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.S3.Key, config.S3.Secret, ""),
		Endpoint:    aws.String(config.S3.Endpoint),
		Region:      aws.String(config.S3.Region),
	})
	if err != nil {
		return err
	}
	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.S3.Bucket),
		Key:    aws.String(url),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return err
	}
	return c.Redirect(urlStr)
}

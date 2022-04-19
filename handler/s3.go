package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/zjyl1994/unilinkd/config"
)

func S3Handler(w http.ResponseWriter, r *http.Request, url string) {
	if config.S3 == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, "S3 not configure")
		return
	}
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.S3.Key, config.S3.Secret, ""),
		Endpoint:    aws.String(config.S3.Endpoint),
		Region:      aws.String(config.S3.Region),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.S3.Bucket),
		Key:    aws.String(url),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	http.Redirect(w, r, urlStr, http.StatusFound)
}

package s3

import (
	"github.com/ameykpatil/go-uploader/constants"
	"github.com/ameykpatil/go-uploader/utils/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"time"
)

var svc *s3.S3

func init() {
	//initialize s3
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(constants.Env.AwsRegion),
		Credentials: credentials.NewStaticCredentials(constants.Env.AwsAccessKeyID, constants.Env.AwsSecretAccessKey, ""),
	})

	// in case of error exit
	if err != nil {
		logger.Err("", "couldn't connect to s3", err)
		os.Exit(1)
	}

	// create S3 service client
	svc = s3.New(sess)
}

//SignedURLForGet returns signed url for getObject for given key
func SignedURLForGet(assetID string, timeout int64) (string, error) {

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(constants.Env.AwsBucket),
		Key:    aws.String(assetID),
	})

	signedURL, err := req.Presign(time.Duration(timeout) * time.Second)

	if err != nil {
		logger.Err(assetID, "failed to get signed url for get object", err)
		return "", err
	}

	logger.Info(assetID, "signedUrl for getObject : "+signedURL)
	return signedURL, nil
}

//SignedURLForPut returns signed url for putObject for given key
func SignedURLForPut(assetID string) (string, error) {

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(constants.Env.AwsBucket),
		Key:    aws.String(assetID),
		//Body:   strings.NewReader(""),
	})

	signedURL, err := req.Presign(15 * time.Minute)

	if err != nil {
		logger.Err(assetID, "failed to get signed url for put object", err)
		return "", err
	}

	logger.Info(assetID, "signedUrl for getObject : "+signedURL)
	return signedURL, nil
}

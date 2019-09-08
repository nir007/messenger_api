package s3

import (
	"io"
	"messenger/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var conf map[interface{}]interface{}

func init() {
	conf, _ = config.Get("s3")
}

// Upload uploads files to object storage
func Upload(file io.Reader, filename string) (string, error) {
	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String(conf["region"].(string)),
		Endpoint: aws.String(conf["endpoint"].(string)),
	})

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(conf["bucket"].(string)),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}

// Download downloads file from storage
func Download() (string, error) {
	/*sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String(conf["region"].(string)),
		Endpoint: aws.String(conf["endpoint"].(string)),
	})

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(conf["bucket"].(string)),
			Key:    aws.String(conf["key"].(string)),
		})

	if err != nil {
		return "", err
	}*/

	return "", nil
}

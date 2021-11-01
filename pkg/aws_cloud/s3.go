package aws_cloud

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var sess = AwsClient()

func UploadToS3(bucket, fileDir string) map[string]string {
	s3Map := make(map[string]string)
	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	pathName := strings.Split(fileDir, "/")
	fileKey := pathName[len(pathName)-1]
	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(fileKey),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	default_url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, fileKey)
	u, _ := uuid.NewV4()
	s3Map["url"] = default_url
	s3Map["id"] = u.String()
	s3Map["bucket"] = bucket
	s3Map["key"] = fileKey

	return s3Map
}

func DownloadFromS3Bucket(bucket, item, path string) map[string]interface{} {
	file, err := os.Create(filepath.Join(path, item))
	if err != nil {
		fmt.Printf("Error in downloading from file: %v \n", err)
		os.Exit(1)
	}

	defer file.Close()

	// Create a downloader with the session and custom options
	downloader := s3manager.NewDownloader(sess, func(d *s3manager.Downloader) {
		d.PartSize = 64 * 1024 * 1024 // 64MB per part
		d.Concurrency = 6
	})

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		fmt.Printf("Error in downloading from file: %v \n", err)
		os.Exit(1)
	}
	resp := make(map[string]interface{})
	resp["name"] = item
	resp["bytes"] = numBytes
	resp["path"] = file.Name()

	return resp
}

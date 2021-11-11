package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	c "saiga/config"
	mo "saiga/models"
	cloud "saiga/pkg/aws_cloud"
	"saiga/pkg/clients"
	"sync"
)

var wg sync.WaitGroup

var path, _ = os.Getwd()

func Initialization() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		InsertDummyData(path + c.Configure().AntonData)
	}()
	go func() {
		defer wg.Done()
		InsertDummyData(path + c.Configure().SteveData)
	}()
	wg.Wait()
}

func InsertDummyData(jsonPath string) {
	jsonFile, err := os.Open(jsonPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s", jsonPath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cm mo.ChatMeta
	json.Unmarshal(byteValue, &cm)

	s3Result := cloud.UploadToS3(c.Configure().S3BucketName, path+c.Configure().DummyImage)

	cm.Messages[0].Attachment.Id = s3Result["id"]
	cm.Messages[0].Attachment.Url = s3Result["url"]
	cm.Messages[0].Attachment.Bucket = s3Result["bucket"]
	cm.Messages[0].Attachment.Key = s3Result["key"]

	_, _ = clients.Mongo.Insert(cm, c.Configure().MongoChatCollection)
}

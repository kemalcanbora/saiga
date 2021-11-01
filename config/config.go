package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	JWTSecret           string `json:"jwt_secret"`
	MongoUserName       string `json:"mongo_user_name"`
	MongoPassword       string `json:"mongo_password"`
	MongoDatabase       string `json:"mongo_database"`
	MongoURL            string `json:"mongo_url"`
	S3BucketName        string `json:"s3_bucket_name"`
	DummyImage          string `json:"dummy_image"`
	DownloadsPath       string `json:"downloads_path"`
	MongoChatCollection string `json:"mongo_chat_collection"`
	AntonData           string `json:"anton_data"`
	SteveData           string `json:"steve_data"`
	AWSAccessKeyId 		string `json:"aws_access_key_id"`
	AWSSecretAccessKey  string `json:"aws_secret_access_key"`
	AWSRegion 			string `json:"aws_region"`
}

func Configure() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var c Config
	c.JWTSecret = os.Getenv("SECRET_KEY")
	c.MongoUserName = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	c.MongoPassword = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	c.MongoDatabase = os.Getenv("MONGO_INITDB_DATABASE")
	c.MongoURL = os.Getenv("MONGO_URL")
	c.S3BucketName = os.Getenv("S3_BUCKET_NAME")
	c.DummyImage = os.Getenv("DUMMY_IMAGE")
	c.DownloadsPath = os.Getenv("DOWNLOADS_PATH")
	c.MongoChatCollection = os.Getenv("MONGO_INITDB_CHAT_COLLECTION")
	c.AntonData = os.Getenv("ANTON_DATA")
	c.SteveData = os.Getenv("STEVE_DATA")
	c.AWSAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	c.AWSSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	c.AWSRegion = os.Getenv("AWS_REGION")
	return c
}

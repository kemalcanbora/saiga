package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	c "saiga/config"
	mo "saiga/models"
	"time"
)

type MongoClient struct {
	Client  *mongo.Client
	Context context.Context
	Cancel  func()
}

func Connection() MongoClient {
	mongoClient := MongoClient{}
	var err error
	credential := options.Credential{
		Username: c.Configure().MongoUserName,
		Password: c.Configure().MongoPassword,
	}
	mongoClient.Client, err = mongo.NewClient(options.Client().ApplyURI(c.Configure().MongoURL).SetAuth(credential))
	mongoClient.Context, mongoClient.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoClient.Cancel()

	err_ := mongoClient.Client.Connect(mongoClient.Context)
	if err_ != nil {
		log.Fatal(err)
	}
	return mongoClient
}

func (m *MongoClient) Insert(data interface{}, collectionName string) (*mongo.InsertOneResult, error) {
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func (m *MongoClient) FindAllTask(collectionName string) []string {
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(collectionName)
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	var tasks []string
	for cur.Next(context.Background()) {
		var result mo.ChatMeta
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, result.ID)
	}
	return tasks
}

func (m *MongoClient) GetChatHistory(collectionName string, Id string) mo.ChatMeta {
	var model mo.ChatMeta
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(collectionName)
	err := collection.FindOne(context.TODO(), bson.M{"_id": Id}).Decode(&model)
	if err != nil {
		fmt.Println(err)
	}
	return model
}

func (m *MongoClient) FindAttachmentURL(collectionName, attachmentId string) (mo.ChatMeta, error) {
	var model mo.ChatMeta
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection(collectionName)

	err := collection.FindOne(context.TODO(),bson.M{"messages.attachment.id": attachmentId}).Decode(&model)
	if err != nil {
		fmt.Println(err)
	}

	return model, err
}

func (m *MongoClient) FindUserWithEmail(user mo.User) mo.User {
	var dbUser mo.User
	collection := m.Client.Database(c.Configure().MongoDatabase).Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		log.Println(err)
	}
	return dbUser
}

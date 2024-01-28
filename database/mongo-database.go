package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amro-alasri/graphQL-server/graph/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "graph"
	COLLECTION = "store"
)

type VideoStore interface {
	Save(Video *model.Video)
	FindAll() []*model.Video
}

type videoStore struct {
	client *mongo.Client
}

func New() VideoStore {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MONGODB := os.Getenv("MONGODB")

	clientOptions := options.Client().ApplyURI(MONGODB)

	clientOptions = clientOptions.SetMaxPoolSize(20)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to MongoDB successfully!")

	return &videoStore{
		client: client,
	}
}

func (s *videoStore) Save(Video *model.Video) {
	collection := s.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), Video)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *videoStore) FindAll() []*model.Video {
	collection := s.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())
	var result []*model.Video

	for cursor.Next(context.TODO()) {
		var v *model.Video
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}

	return result
}

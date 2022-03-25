// Package repository  for db stores
package repository

import (
	"context"
	"fmt"
	"github.com/ahmedkhaeld/graphql-server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// VideoRepository for video repo
type VideoRepository interface {
	Save(video *model.Video)
	FindAll() []*model.Video
}

//database to implement VideoRepository interface
type database struct {
	client *mongo.Client
}

const (
	DATABASE   = "graphql"
	COLLECTION = "videos"
)

// NewVideoRepository initialize the repo, to connect the app with mongodb
func NewVideoRepository() VideoRepository {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	MONGODB := os.Getenv("MONGODB")

	// Set client options
	clientOptions := options.Client().ApplyURI(MONGODB)

	clientOptions = clientOptions.SetMaxPoolSize(50)

	// Connect to MongoDB
	userClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = userClient.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &database{
		client: userClient,
	}
}

// Save implementation to save a new video to the store
func (db *database) Save(video *model.Video) {
	// access the database graph collection videos
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	// store video to the collection
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		log.Fatal(err)
	}
}

// FindAll implementation to fetch all videos from the store
func (db *database) FindAll() []*model.Video {
	// access the database graph collection videos
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	// user the Find function from the collection to get all the videos with no filter
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	// close the cursor once the function finishes
	defer cursor.Close(context.TODO())
	// create a var to store the results from the db
	var results []*model.Video
	// loop through documents using Next func
	for cursor.Next(context.TODO()) {
		var v *model.Video
		// decode each element
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		// append each video to the results slice
		results = append(results, v)
	}
	return results
}

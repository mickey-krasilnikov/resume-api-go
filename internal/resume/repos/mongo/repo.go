package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database   string
	collection string
)

const (
	// environment variables
	mongoDBConnectionStringEnvVarName = "MONGODB_CONNECTION_STRING"
	mongoDBDatabaseEnvVarName         = "MONGODB_DATABASE"
	mongoDBCollectionEnvVarName       = "MONGODB_COLLECTION"
)

type MongoResumeRepository struct {
}

func New() *MongoResumeRepository {
	return &MongoResumeRepository{}
}

// connects to MongoDB
func connect() *mongo.Client {
	mongoDBConnectionString := os.Getenv(mongoDBConnectionStringEnvVarName)
	if mongoDBConnectionString == "" {
		log.Fatal("missing environment variable: ", mongoDBConnectionStringEnvVarName)
	}

	database = os.Getenv(mongoDBDatabaseEnvVarName)
	if database == "" {
		log.Fatal("missing environment variable: ", mongoDBDatabaseEnvVarName)
	}

	collection = os.Getenv(mongoDBCollectionEnvVarName)
	if collection == "" {
		log.Fatal("missing environment variable: ", mongoDBCollectionEnvVarName)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoDBConnectionString).SetDirect(true)
	c, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Printf("unable to create client %v", err)
	}
	err = c.Connect(ctx)

	if err != nil {
		log.Printf("unable to initialize connection %v", err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Printf("unable to connect %v", err)
	}
	return c
}

// creates a resume
func (repo *MongoResumeRepository) CreateResume(resume models.Resume) error {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	resumeCollection := c.Database(database).Collection(collection)
	r, err := resumeCollection.InsertOne(ctx, resume)
	if err != nil {
		log.Printf("failed to add resume %v", err)
	}
	fmt.Println("added resume", r.InsertedID)
	return err
}

// lists resumes
func (repo *MongoResumeRepository) GetAllResumes() ([]models.Resume, error) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	resumeCollection := c.Database(database).Collection(collection)
	rs, err := resumeCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("failed to list resume(s) %v", err)
	}
	var resumes []models.Resume
	err = rs.All(ctx, &resumes)
	if err != nil {
		log.Printf("failed to list resume(s) %v", err)
	}

	return resumes, err
}

// get resume by id
func (repo *MongoResumeRepository) GetResumeById(resumeid string) (models.Resume, error) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	resumeCollection := c.Database(database).Collection(collection)
	oid, err := primitive.ObjectIDFromHex(resumeid)
	if err != nil {
		log.Printf("failed to update resume %v", err)
		return models.Resume{}, err
	}
	rs := resumeCollection.FindOne(ctx, bson.D{{Key: "_id", Value: oid}})

	var resume models.Resume
	err = rs.Decode(&resume)
	if err != nil {
		log.Printf("failed to list resume(s) %v", err)
		return models.Resume{}, err
	}

	return resume, nil
}

// updates resume
func (repo *MongoResumeRepository) UpdateResume(resumeid string, newResume models.Resume) error {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	resumeCollection := c.Database(database).Collection(collection)
	oid, err := primitive.ObjectIDFromHex(resumeid)
	if err != nil {
		log.Printf("failed to update resume %v", err)
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	_, err = resumeCollection.ReplaceOne(ctx, filter, newResume)
	if err != nil {
		log.Printf("failed to update resume %v", err)
		return err
	}
	return nil
}

// deletes a resume
func (repo *MongoResumeRepository) DeleteResume(resumeid string) error {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	resumeCollection := c.Database(database).Collection(collection)
	oid, err := primitive.ObjectIDFromHex(resumeid)
	if err != nil {
		log.Printf("invalid resume ID %v", err)
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	_, err = resumeCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("failed to delete resume %v", err)
		return err
	}
	return nil
}

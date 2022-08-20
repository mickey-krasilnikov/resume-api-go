package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
	"go.mongodb.org/mongo-driver/bson"
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
	err = c.Connect(ctx)

	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
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
		log.Fatalf("failed to add resume %v", err)
	}
	fmt.Println("added resume", r.InsertedID)
	return err
}

// lists resumes
func (repo *MongoResumeRepository) ListResumes() ([]models.Resume, error) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	resumeCollection := c.Database(database).Collection(collection)
	rs, err := resumeCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatalf("failed to list resume(s) %v", err)
	}
	var resumes []models.Resume
	err = rs.All(ctx, &resumes)
	if err != nil {
		log.Fatalf("failed to list resume(s) %v", err)
	}

	return resumes, err
}

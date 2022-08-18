package mongo

import (
	"context"
	"os"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ResumeCollection = "resume"
)

type MongoResumeRepository struct {
	db *mongo.Database
}

func New() (*MongoResumeRepository, error) {
	var dbName = "resume"
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	if err = client.Connect(context.TODO()); err != nil {
		panic(err)
	}
	db := client.Database(dbName)
	//db.CreateCollection(context.Background(), ResumeCollection)
	return &MongoResumeRepository{db: db}, err
}

func (r *MongoResumeRepository) GetResume(ctx context.Context) (*models.Resume, error) {
	resume := new(models.Resume)
	res := r.db.Collection(ResumeCollection).FindOne(ctx, bson.M{"_id": 0})
	if err := res.Err(); err != nil {
		return nil, err
	}
	err := res.Decode(resume)
	return resume, err
}

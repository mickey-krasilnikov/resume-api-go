package memory

import (
	"errors"
	"log"
	"time"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ResumeNotFound = errors.New("Resume Not Found")

type InMemoryResumeRepo struct {
	resumes []models.Resume
}

func New() *InMemoryResumeRepo {
	return &InMemoryResumeRepo{resumes: []models.Resume{models.ResumeMock}}
}

func (repo *InMemoryResumeRepo) GetAllResumes() ([]models.Resume, error) {
	return repo.resumes, nil
}

func (repo *InMemoryResumeRepo) GetResumeById(resumeid string) (models.Resume, error) {
	oid, err := primitive.ObjectIDFromHex(resumeid)
	if err != nil {
		log.Printf("failed to update resume %v", err)
	}
	for _, s := range repo.resumes {
		if s.ID == oid {
			return s, nil
		}
	}
	return models.Resume{}, ResumeNotFound
}

func (repo *InMemoryResumeRepo) CreateResume(resume models.Resume) error {
	resume.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	repo.resumes = append(repo.resumes, resume)
	return nil
}

func (repo *InMemoryResumeRepo) DeleteResume(resumeid string) error {
	oid, err := primitive.ObjectIDFromHex(resumeid)
	if err != nil {
		log.Printf("failed to update resume %v", err)
	}
	for i, s := range repo.resumes {
		if s.ID == oid {
			repo.resumes[i] = repo.resumes[len(repo.resumes)-1]
			repo.resumes = repo.resumes[:len(repo.resumes)-1]
			return nil
		}
	}
	return ResumeNotFound
}

func (repo *InMemoryResumeRepo) UpdateResume(resumeid string, newResume models.Resume) error {
	oid, err := primitive.ObjectIDFromHex(resumeid)
	if err != nil {
		log.Printf("failed to update resume %v", err)
		return err
	}
	for i, s := range repo.resumes {
		if s.ID == oid {
			newResume.ID = oid
			repo.resumes[i] = newResume
			break
		}
	}
	return nil
}

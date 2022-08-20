package memory

import (
	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
)

type InMemoryResumeRepo struct {
	resumes []models.Resume
}

func New() *InMemoryResumeRepo {
	return &InMemoryResumeRepo{resumes: []models.Resume{models.ResumeMock}}
}

func (repo *InMemoryResumeRepo) ListResumes() ([]models.Resume, error) {
	resumes := []models.Resume{models.ResumeMock}
	return resumes, nil
}

func (repo *InMemoryResumeRepo) CreateResume(resume models.Resume) error {
	repo.resumes = append(repo.resumes, resume)
	return nil
}

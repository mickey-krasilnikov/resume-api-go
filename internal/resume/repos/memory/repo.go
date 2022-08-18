package memory

import (
	"context"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
)

type InMemoryResumeRepo struct {
	resume *models.Resume
}

func New() (*InMemoryResumeRepo, error) {
	return &InMemoryResumeRepo{resume: &ResumeMock}, nil
}

func (r *InMemoryResumeRepo) GetResume(ctx context.Context) (*models.Resume, error) {
	return r.resume, nil
}

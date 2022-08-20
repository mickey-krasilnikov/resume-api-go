package repos

import (
	"errors"
	"fmt"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
	memrepo "github.com/mickey-krasilnikov/resume-api-go/internal/resume/repos/memory"
	mongorepo "github.com/mickey-krasilnikov/resume-api-go/internal/resume/repos/mongo"
	//sqlrepo "github.com/mickey-krasilnikov/resume-api-go/internal/resume/repos/sql"
)

type ResumeRepo interface {
	CreateResume(resume models.Resume) error
	ListResumes() ([]models.Resume, error)
}

type RepoType string

const (
	InMemory RepoType = "InMemory"
	//SQL      RepoType = "SQL"
	Mongo RepoType = "Mongo"
)

func GetRepo(t RepoType) (ResumeRepo, error) {
	var repo ResumeRepo
	var err error
	switch t {
	case InMemory:
		repo = memrepo.New()
	//case SQL:
	//	repo, err = sqlrepo.New()
	case Mongo:
		repo = mongorepo.New()
	default:
		repo = nil
		err = errors.New(fmt.Sprintf("%s repository type is not supported", t))
	}
	return repo, err
}

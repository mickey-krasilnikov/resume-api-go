package sql

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/lib/pq"
// 	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
// )

// const (
// 	ResumeTable = "resume"
// )

// type SqlResumeRepository struct {
// 	db *sqlx.DB
// }

// func New() (*SqlResumeRepository, error) {
// 	db, err := sqlx.Connect("postgres", os.Getenv("PG_URI"))
// 	return &SqlResumeRepository{db: db}, err
// }

// func (r *SqlResumeRepository) GetResume(ctx context.Context) (*models.Resume, error) {
// 	p := new(models.Resume)
// 	sql := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", ResumeTable)
// 	err := r.db.GetContext(ctx, p, sql, 0)
// 	return p, err
// }

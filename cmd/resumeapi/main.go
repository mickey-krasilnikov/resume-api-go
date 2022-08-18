package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/repos"
)

func main() {
	repo, err := repos.GetRepo(repos.InMemory)
	if err != nil {
		panic(fmt.Sprintf("unable to create repo. application crashed with an error:\n%s", err))
	}
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8081")
	}
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	registerRoutes(r, repo)
	r.Run()
}

func registerRoutes(r *gin.Engine, repo repos.ResumeRepo) {
	g := r.Group("/api")
	g.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/api/resume")
	})
	g.GET("/resume", func(ctx *gin.Context) {
		resume, err := repo.GetResume(context.Background())
		if err != nil {
			err := &gin.Error{
				Err:  err,
				Type: gin.ErrorTypeRender,
				Meta: "error while getting resume from repo",
			}
			ctx.Error(err)
		} else {
			ctx.JSON(http.StatusOK, resume)
			return
		}
	})
}

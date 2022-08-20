package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/models"
	"github.com/mickey-krasilnikov/resume-api-go/internal/resume/repos"
)

func main() {
	repo, err := repos.GetRepo(repos.Mongo)
	if err != nil {
		panic(fmt.Sprintf("unable to access the repo. application crashed with an error:\n%s", err))
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
	g.GET("/resumes", func(ctx *gin.Context) {
		resumes, err := repo.ListResumes()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, resumes)
	})

	g.POST("/resumes", func(ctx *gin.Context) {
		var resume models.Resume
		if err := ctx.ShouldBind(&resume); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		if err := repo.CreateResume(resume); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusCreated, nil)
	})
}

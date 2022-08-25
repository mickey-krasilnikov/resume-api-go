package main

import (
	"log"
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
		log.Fatalf("unable to access the repo. application crashed with an error:\n%s", err)
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
	g := r.Group("/api/resumes")

	g.GET("/", func(ctx *gin.Context) {
		resumes, err := repo.GetAllResumes()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, resumes)
	})
	g.GET("/:resumeID", func(ctx *gin.Context) {
		resumeID := ctx.Param("resumeID")
		resume, err := repo.GetResumeById(resumeID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}
		ctx.JSON(http.StatusOK, resume)

	})
	g.POST("/", func(ctx *gin.Context) {
		var resume models.Resume
		if err := ctx.ShouldBind(&resume); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		if err := repo.CreateResume(resume); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusCreated, nil)
	})
	g.PUT("/:resumeID", func(ctx *gin.Context) {
		resumeID := ctx.Param("resumeID")
		var resume models.Resume
		if err := ctx.ShouldBind(&resume); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		if err := repo.UpdateResume(resumeID, resume); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, nil)
	})
	g.DELETE("/:resumeID", func(ctx *gin.Context) {
		resumeID := ctx.Param("resumeID")
		if err := repo.DeleteResume(resumeID); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, nil)
	})
}

package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/mickey-krasilnikov/resume-api-go/resume"
)

func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8081")
	}
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	registerRoutes(r)
	r.Run()
}

func registerRoutes(r *gin.Engine) {
	g := r.Group("/api")
	g.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/api/resume")
	})
	g.GET("/resume", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, resume.ResumeMock)
	})
	g.GET("/resume/contacts", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, resume.ResumeMock.Contacts)
	})
	g.GET("/resume/summary", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, resume.ResumeMock.Summary)
	})
	g.GET("/resume/skills", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, resume.ResumeMock.Skills)
	})
	g.GET("/resume/experience", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, resume.ResumeMock.Experience)
	})
	g.GET("/resume/certifications", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, resume.ResumeMock.Certifications)
	})
	g.GET("/resume/education", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, resume.ResumeMock.Education)
	})
}

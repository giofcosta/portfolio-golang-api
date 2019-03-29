package controllers

import (
	"net/http"

	"github.com/giofcosta/portfolio-golang-api/repositories"

	"github.com/gin-gonic/gin"
)

type ResumeController struct {
	Repository repositories.ResumeRepository
}

func (r *ResumeController) GET(c *gin.Context) {
	c.JSON(http.StatusOK, r.Repository.GetAll())
}

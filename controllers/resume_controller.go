package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giofcosta/portfolio-golang-api/models"
)

type ResumeController struct {
}

func (r *ResumeController) GET(c *gin.Context) {
	model := &models.Resume{}
	c.JSON(http.StatusOK, model.GetAll())
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giofcosta/portfolio-golang-api/models"
)

//SkillController has the data structure of skills
type SkillController struct{}

//GET returns the skills
func (s SkillController) GET(c *gin.Context) {
	model := &models.Skill{}
	c.JSON(http.StatusOK, model.GetAll())
}

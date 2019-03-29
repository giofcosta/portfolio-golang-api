package controllers

import (
	"net/http"

	"github.com/giofcosta/portfolio-golang-api/repositories"

	"github.com/gin-gonic/gin"
)

//SkillController has the data structure of skills
type SkillController struct {
	Repository repositories.SkillRepository
}

//GET returns the skills
func (s SkillController) GET(c *gin.Context) {
	c.JSON(http.StatusOK, s.Repository.GetAll())
}

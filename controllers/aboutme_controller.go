package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giofcosta/portfolio-golang-api/models"
)

//AboutMeController controls the about me page informations
type AboutMeController struct{}

//GET returns the about me page information
func (am AboutMeController) GET(c *gin.Context) {
	model := &models.AboutMe{}
	c.JSON(http.StatusOK, model.GetAll())
}

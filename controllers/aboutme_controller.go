package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giofcosta/portfolio-golang-api/repositories"
)

//AboutMeController controls the about me page informations
type AboutMeController struct {
	Repository repositories.AboutMeRepository
}

//GET returns the about me page information
func (a AboutMeController) GET(c *gin.Context) {
	c.JSON(http.StatusOK, a.Repository.GetAll())
}

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/giofcosta/portfolio-golang-api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	skill := new(controllers.SkillController)
	aboutme := new(controllers.AboutMeController)
	resume := new(controllers.ResumeController)

	router.GET("/skill", skill.GET)
	router.GET("/aboutme", aboutme.GET)
	router.GET("/resume", resume.GET)

	return router
}

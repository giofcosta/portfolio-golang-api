package server

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(app *Application) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/skill", app.SkillController.GET)
	router.GET("/aboutme", app.AboutMeController.GET)
	router.GET("/resume", app.ResumeController.GET)

	return router
}

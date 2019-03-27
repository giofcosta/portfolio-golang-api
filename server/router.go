package server

import (
	"github.com/gin-gonic/gin"
	"github.com/giofcosta/portfolio-golang-api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	skill := new(controllers.SkillController)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/skill", skill.GET)

	return router
}

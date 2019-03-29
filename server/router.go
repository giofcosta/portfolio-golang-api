package server

import (
	"fmt"
	"log"
	"os"

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

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	docDir := dir + "/docs/doc.html"

	if _, err := os.Stat(docDir); err == nil {
		fmt.Println(dir+"/docs/doc.html", "path/to/whatever exists")
		router.LoadHTMLFiles(docDir)

		router.GET("/v1/doc", func(c *gin.Context) {
			c.HTML(200, "doc.html", gin.H{})
		})
	} else {
		log.Fatal("path to doc does *not* exist")
	}

	router.GET("/v1/skill", app.SkillController.GET)
	router.GET("/v1/aboutme", app.AboutMeController.GET)
	router.GET("/v1/resume", app.ResumeController.GET)

	return router
}

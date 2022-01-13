package main

import (
	"go_api/controller"
	"go_api/middlewares"
	"go_api/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func Home(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()
	server := gin.New()

	server.Static("/css", "./template/css")

	server.LoadHTMLGlob("templates/*.html")

	server.SetTrustedProxies([]string{"localhost"})

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")

	server.GET("/", Home)

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(c *gin.Context) {
		err := videoController.Save(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Video input is valid",
			})
		}
		c.JSON(200, videoController.Save(c))
	})

	server.Run(":4444")
}

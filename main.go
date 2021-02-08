package main

import (
	"ginrest/gin-rest-middleware-mongo/src/controller"
	"ginrest/gin-rest-middleware-mongo/src/middlewares"
	"ginrest/gin-rest-middleware-mongo/src/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	profileService    service.ProfileService       = service.New()
	profileController controller.ProfileController = controller.New(profileService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())
	r.GET("/profile", func(c *gin.Context) {
		c.JSON(200, profileController.FindAll())
	})
	r.POST("/profile", func(c *gin.Context) {
		c.JSON(200, profileController.Save(c))
	})
	r.Run(":8081")
}

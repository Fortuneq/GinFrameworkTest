package main

import (
	"ginframe/controller"
	"ginframe/service"

	"github.com/gin-gonic/gin"
)


var (
	videoService service.VideoService = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)
func main(){

	server := gin.Default()

	server.GET("/post",func(ctx *gin.Context) {
		ctx.JSON(200,VideoController.FindAll())
	})

	server.POST("/post",func(ctx *gin.Context) {
		ctx.JSON(200,VideoController.Save(ctx))
	})
	server.Run(":8080")
}
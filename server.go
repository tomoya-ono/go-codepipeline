package main

import (
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"gin-test/service"
	"gin-test/controller"
	"gin-test/middlewares"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)



func main() {
	// server := gin.Default()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context){
			ctx.JSON(200, videoController.FindAll())
		})
	
		apiRoutes.POST("/videos", func(ctx *gin.Context){
			// ctx.JSON(200, videoController.Save(ctx))
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}else{
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is valid"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)

}
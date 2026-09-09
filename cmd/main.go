package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrandrey24/url-shortening-server/internal/controller"
	local "github.com/mrandrey24/url-shortening-server/internal/data"
	"github.com/mrandrey24/url-shortening-server/internal/service"
)

func main() {

	repo := local.NewMemoryURLRepository()
	src := service.NewURLService(repo)
	ctrl := controller.NewURLController(src)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		shorter := v1.Group("/shorter")
		{
			const codePath = "/:code"
			shorter.POST("", ctrl.CreateShortUrl)
			shorter.PUT(codePath, ctrl.UpdateShortUrl)
			shorter.GET(codePath, ctrl.GetShortUrlByCode)
			shorter.DELETE(codePath, ctrl.DeleteShortUrl)

		}
	}

	router.Run(":8080")
}

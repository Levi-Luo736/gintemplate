package router

import (
	"embed"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/levi/gin-template/common"
	"github.com/levi/gin-template/controller"
	"github.com/levi/gin-template/middleware"
)

func setWebRouter(router *gin.Engine, buildFS embed.FS, indexPage []byte) {
	router.Use(middleware.GlobalWebRateLimit())
	fileDownloadRoute := router.Group("/")
	fileDownloadRoute.GET("/upload/:file", middleware.DownloadRateLimit(), controller.DownloadFile)
	router.Use(middleware.Cache())
	router.Use(static.Serve("/", common.EmbedFolder(buildFS, "web/build")))
	router.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})
}

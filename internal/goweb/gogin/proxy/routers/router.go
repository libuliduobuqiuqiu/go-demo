package routers

import (
	"fmt"
	"godemo/internal/goweb/gogin/proxy/public"
	"net/http"

	"github.com/gin-gonic/gin"
)

var registerFunc []func(*gin.RouterGroup)

func StartProxyService(address string, port int) {
	router := gin.New()
	router.Use(gin.Recovery())

	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, public.ProxyResponse{
			Err:     http.StatusNotFound,
			Message: "The request url was not found on the server",
		})
	})

	// Register router
	rootGroup := router.Group("netac/base")
	for _, f := range registerFunc {
		f(rootGroup)
	}

	router.Run(fmt.Sprintf("%s:%d", address, port))
}

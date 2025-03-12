package routers

import (
	"godemo/internal/goweb/gogin/proxy/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	registerFunc = append(registerFunc, RegisterNetworkRouters)
}

func RegisterNetworkRouters(group *gin.RouterGroup) {
	networkGroup := group.Group("")
	networkGroup.GET("terminal", handlers.Terminal)
}

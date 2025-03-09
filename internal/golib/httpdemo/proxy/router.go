package proxy

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitProxyRouter(address string, port int) {
	router := gin.New()
	router.Use(gin.Recovery())

	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	group := router.Group("netac/base")
	group.Any("proxy", ProxyHttpReq)
	group.POST("terminal", ProxyTerminalReq)
	group.POST("ssh", ProxySshReq)

	router.Run(fmt.Sprintf("%s:%d", address, port))
}

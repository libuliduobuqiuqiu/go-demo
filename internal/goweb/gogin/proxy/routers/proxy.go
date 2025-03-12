package routers

import "github.com/gin-gonic/gin"

func init() {
	registerFunc = append(registerFunc, RegisterProxyRouter)
}

func RegisterProxyRouter(group *gin.RouterGroup) {
	proxyGroup := group.Group("proxy")
	proxyGroup.Any("")
}

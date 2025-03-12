package routers

import (
	"godemo/internal/goweb/gogin/proxy/handlers"

	"github.com/gin-gonic/gin"
)

func init() {
	registerFunc = append(registerFunc, RegisterSshRouters)
}

func RegisterSshRouters(group *gin.RouterGroup) {
	sshGroup := group.Group("")
	sshGroup.POST("ssh", handlers.ExecSshCommand)
}

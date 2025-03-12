package handlers

import (
	"errors"
	"godemo/internal/goweb/gogin/proxy/public"
	"godemo/internal/goweb/gogin/proxy/service"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	address := c.Query("address")
	timeout := c.DefaultQuery("timeout", "30")
	count := c.DefaultQuery("count", "5")

	if address == "" {
		err := errors.New("address 参数不能为空")
		public.HandleErrJson(c, err)
		return
	}

	networkService := service.NewNetworkService()
	res, err := networkService.Ping(address, timeout, count)
	if err != nil {
		public.HandleErrJson(c, err)
		return
	}

	resp := struct {
		PingRes string `json:"ping_res"`
	}{
		PingRes: res,
	}
	public.HandleSuccessJson(c, resp)
}

func TraceRoute(c *gin.Context) {

}

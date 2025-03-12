package handlers

import (
	"errors"
	"fmt"
	"godemo/internal/goweb/gogin/proxy/public"
	"godemo/internal/goweb/gogin/proxy/service"
	"strconv"

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
	address := c.Query("address")
	if address == "" {
		public.HandleErrJson(c, errors.New("address param is not empty."))
		return
	}

	tmpCount := c.DefaultQuery("count", "30")
	tmpTimeout := c.DefaultQuery("timeout", "300")
	tmpTtl := c.DefaultQuery("max_ttl", "30")

	count, timeout, ttl, err := checkTraceRouteParam(tmpCount, tmpTimeout, tmpTtl)
	if err != nil {
		public.HandleErrJson(c, err)
		return
	}

	networkService := service.NewNetworkService()
	res, err := networkService.Traceroute(address, count, timeout, ttl)
	if err != nil {
		public.HandleErrJson(c, err)
		return
	}
	resp := struct {
		TracerouteRes []string `json:"traecroute_res"`
	}{
		TracerouteRes: res,
	}

	public.HandleSuccessJson(c, resp)
	return
}

func checkTraceRouteParam(count, timeout, ttl string) (c, t, tt int, err error) {
	c, err = strconv.Atoi(count)
	if err != nil {
		err = errors.New(fmt.Sprintf("count param error: %s", err.Error()))
		return
	}

	t, err = strconv.Atoi(timeout)
	if err != nil {
		err = errors.New(fmt.Sprintf("timeout param error: %s", err.Error()))
		return
	}

	tt, err = strconv.Atoi(ttl)
	if err != nil {
		err = errors.New(fmt.Sprintf("ttl param error: %s", err.Error()))
		return
	}

	return
}

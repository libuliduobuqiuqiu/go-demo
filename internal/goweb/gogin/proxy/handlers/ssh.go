package handlers

import (
	"net/http"

	"godemo/internal/goweb/gogin/proxy/public"
	"godemo/internal/goweb/gogin/proxy/service"

	"github.com/gin-gonic/gin"
)

func ExecSshCommand(ctx *gin.Context) {
	var (
		params public.ProxySshParams
		resp   public.ProxyResponse
		cmdRes map[string]string
	)

	if err := ctx.BindJSON(&params); err != nil {
		public.HandleErrJson(ctx, err)
		return
	}

	sshService := service.NewSshService()
	cmdRes, err := sshService.ExecCommand(params)
	if err != nil {
		public.HandleErrJson(ctx, err)
	}

	resp.Data = cmdRes
	ctx.JSON(http.StatusOK, resp)
}

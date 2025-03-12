package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"godemo/internal/goweb/gogin/proxy/public"
	"golang.org/x/crypto/ssh"
)

func SshConnect(address, username, password string, port int, isKeyboardInteractive bool) (client *ssh.Client, err error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if isKeyboardInteractive {
		config.Auth = []ssh.AuthMethod{setKeyboardInteractive(password)}
	}

	if port != 0 {
		address = fmt.Sprintf("%s:%d", address, port)
	} else {
		address = fmt.Sprintf("%s:22", address)
	}

	client, err = ssh.Dial("tcp", address, config)
	if err != nil {
		return
	}

	return
}

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

	client, err := SshConnect(params.Address, params.Username, params.Password, params.Port, true)
	if err != nil {
		public.HandleErrJson(ctx, err)
		return
	}

	defer client.Close()
	cmdRes = make(map[string]string)
	for _, cmd := range params.Commands {
		session, err := client.NewSession()
		if err != nil {
			public.HandleErrJson(ctx, err)
			return
		}
		defer session.Close()
		var stdoutBuf, stderrBuf bytes.Buffer
		session.Stdout = &stdoutBuf
		session.Stderr = &stderrBuf
		if err = session.Run(cmd); err != nil {
			err = fmt.Errorf("Run command:%s:%s, error: %w", cmd, stderrBuf.String(), err)
			public.HandleErrJson(ctx, err)
			return
		}
		cmdRes[cmd] = stdoutBuf.String()
	}

	resp.Data = cmdRes
	ctx.JSON(http.StatusOK, resp)
}

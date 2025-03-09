package proxy

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
)

func ProxySshReq(ctx *gin.Context) {
	var (
		address string
		params  ProxySshParams
		resp    ProxyResponse
		cmdRes  map[string]string
	)

	if err := ctx.BindJSON(&params); err != nil {
		log.Fatal(err)
	}

	config := &ssh.ClientConfig{
		User: params.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(params.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if params.Port != 0 {
		address = fmt.Sprintf("%s:%d", params.Address, params.Port)
	} else {
		address = params.Address
	}

	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	cmdRes = make(map[string]string)
	for _, cmd := range params.Commands {
		session, err := client.NewSession()
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()

		var stdoutBuf, stderrBuf bytes.Buffer
		session.Stdout = &stdoutBuf
		session.Stderr = &stderrBuf
		if err = session.Run(cmd); err != nil {
			err = fmt.Errorf("Run command:%s:%s, error: %w", cmd, stderrBuf.String(), err)
			log.Fatal(err)
		}
		cmdRes[cmd] = stdoutBuf.String()
	}

	resp.Data = cmdRes
	ctx.JSON(http.StatusOK, resp)
}

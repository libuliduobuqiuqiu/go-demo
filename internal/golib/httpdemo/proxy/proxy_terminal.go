package proxy

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
)

// 1. 接受http请求，升级为ws请求
// 2. 根据传入的http请求参数，新建ssh配置，初始化client，根据client新建session
// 3. 发送建立伪终端, 建立通道将终端输出到ws中，将ws读取的命令输入到终端

func ProxyTerminalReq(ctx *gin.Context) {
	// 升级为Websocket请求
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	r := ctx.Request
	w := ctx.Writer
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		HandleErrJson(w, err)
		return
	}
	defer ws.Close()

	// 读取传入参数
	params := ProxyParams{}
	if err = ws.ReadJSON(&params); err != nil {
		HandleErrMessage(ws, err)
		return
	}

	// 打开Ssh Client
	client, err := NewSshClient(params.Address, params.Username, params.Password)
	if err != nil {
		HandleErrMessage(ws, err)
		return
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		HandleErrMessage(ws, err)
		return
	}

	if err = session.RequestPty("xterm", 768, 1024, ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}); err != nil {
		HandleErrMessage(ws, err)
		return
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		HandleErrMessage(ws, err)
		return
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		HandleErrMessage(ws, err)
		return
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		HandleErrMessage(ws, err)
		return
	}

	go io.Copy(wsWriter(ws), stdout)
	go io.Copy(wsWriter(ws), stderr)
	go io.Copy(stdin, wsReader(ws))

	if err := session.Shell(); err != nil {
		HandleErrMessage(ws, err)
		return
	}

	if err := session.Wait(); err != nil {
		log.Printf("session ended with error: %v", err)
	}

}

func NewSshClient(address, username, password string) (client *ssh.Client, err error) {
	clientConfig := ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			setKeyboardInteractive(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err = ssh.Dial("tcp", address, &clientConfig)
	if err != nil {
		return
	}

	return
}

func wsReader(ws *websocket.Conn) io.Reader {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}
			pw.Write(msg)
		}
	}()

	return pr
}

func wsWriter(ws *websocket.Conn) io.Writer {
	pr, pw := io.Pipe()
	go func() {
		defer pr.Close()
		buf := make([]byte, 1024)
		for {
			n, err := pr.Read(buf)
			if err != nil {
				return
			}

			ws.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()
	return pw
}

// 设置键盘交互
func setKeyboardInteractive(password string) ssh.AuthMethod {
	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {
		if len(questions) == 0 {
			return []string{}, nil
		}
		return []string{password}, nil
	}
	return ssh.KeyboardInteractive(keyboardInteractiveChallenge)
}

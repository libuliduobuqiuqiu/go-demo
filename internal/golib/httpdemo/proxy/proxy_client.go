package proxy

import (
	"bufio"
	"context"
	"fmt"
	"godemo/pkg"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func CommitDeviceHttpReq(reqUrl string) (err error) {
	client := http.Client{Timeout: 60 * time.Second}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, reqUrl, nil)
	if err != nil {
		return
	}

	req.SetBasicAuth("admin", "admin")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	raws, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(raws))
	return nil
}

func CommitDeviceSshReq(serviceAddress string) {
	// 建立websocket连接
	endpoint := "/netac/base/ssh"
	reqUrl := url.URL{Scheme: "ws", Host: serviceAddress, Path: endpoint}
	ws, resp, err := websocket.DefaultDialer.Dial(reqUrl.String(), nil)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		log.Printf(string(body))
		log.Printf(err.Error())
		return
	}

	config := pkg.GetGlobalConfig("")

	params := ProxyParams{
		Username: config.F5Config.Username,
		Password: config.F5Config.Password,
		Address:  fmt.Sprintf("%s:%d", config.F5Config.Host, config.F5Config.Port),
	}

	if err := ws.WriteJSON(params); err != nil {
		log.Printf(err.Error())
		return
	}

	// 启动前台控制端
	go ProxyShell(ws)

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		fmt.Printf(string(message))
	}
}

// 读取输入
func ProxyShell(ws *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text() + "\n"
		if err := ws.WriteMessage(websocket.TextMessage, []byte(command)); err != nil {
			log.Printf(err.Error())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf(err.Error())
	}
}

package proxy

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ProxyResponse struct {
	Err     int         `json:"err"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ProxySshParams struct {
	ProxyParams
	Commands []string `json:"commands"`
}

type ProxyParams struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port,omitempty"`
}

func genErrInfo(err error) []byte {
	log.Printf(err.Error())
	r := ProxyResponse{
		Err:     500,
		Message: err.Error(),
	}

	resp, err := json.Marshal(r)
	if err != nil {
		log.Printf("Marshal Error.")
	}
	return resp
}

func HandleErrJson(ctx *gin.Context, err error) {
	logrus.Warn(err)
	resp := ProxyResponse{
		Err:     500,
		Message: err.Error(),
	}

	ctx.JSON(http.StatusInternalServerError, resp)
}

func HandleErrMessage(ws *websocket.Conn, err error) {
	resp := genErrInfo(err)
	ws.WriteMessage(websocket.TextMessage, resp)
}

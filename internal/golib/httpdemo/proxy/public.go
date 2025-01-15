package proxy

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type ProxyResponse struct {
	Err     int         `json:"err`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ProxyParams struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
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

func HandleErrJson(w http.ResponseWriter, err error) {
	resp := genErrInfo(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(resp)
}

func HandleErrMessage(ws *websocket.Conn, err error) {
	resp := genErrInfo(err)
	ws.WriteMessage(websocket.TextMessage, resp)
}

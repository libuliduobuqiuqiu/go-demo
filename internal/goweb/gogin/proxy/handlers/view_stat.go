package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
)

// view stat api (generate fake data)

type ViewStatHandlers struct{}

func NewViewStatHandlers() *ViewStatHandlers {
	return &ViewStatHandlers{}
}

type QueryParam struct {
	SecretKey   string   `json:"secret_key"`
	ProbeIds    string   `json:"probeIds"`
	Date        string   `json:"date"`
	Step        string   `json:"step"`
	ReturnField []string `json:"returnField"`
	TlvFilter   string   `json:"tlvFilter"`
}

type ViewStat struct {
	Serveripaddr       string `json:"serveripaddr"`
	ServerPort         string `json:"serverPort"`
	BytesIn            int64  `json:"bytesIn"`
	BytesOut           int64  `json:"bytesOut"`
	NewConnections     int64  `json:"newConnections"`
	CurrentConnections int64  `json:"currentConnections"`
	TotalBitps         int64  `json:"totalBitps"`
	ResponseTime       int64  `json:"responseTime"`
}

func (v *ViewStatHandlers) GetData(c *gin.Context) {
	tmpQuery := c.Query("query")
	if tmpQuery == "" {
		err := errors.New("query参数不能为空")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(c.Request.URL)

	fmt.Println(url.PathUnescape(c.Request.URL.String()))

	fmt.Println(tmpQuery)

	var query QueryParam
	if err := json.Unmarshal([]byte(tmpQuery), &query); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var resp []*ViewStat
	tlvFilter := query.TlvFilter
	reg := regexp.MustCompile(`\(serverIpAddr=([\w\.]+)&&serverPort=([\w\.]+)\)`)
	filters := reg.FindAllStringSubmatch(tlvFilter, -1)
	for _, f := range filters {
		addr := f[1]
		port := f[2]
		fmt.Println(f)

		v := &ViewStat{
			Serveripaddr:       addr,
			ServerPort:         port,
			BytesIn:            1000,
			BytesOut:           2000,
			NewConnections:     200,
			CurrentConnections: 100,
			TotalBitps:         3024,
			ResponseTime:       10,
		}
		resp = append(resp, v)
	}

	respData := struct {
		Data []*ViewStat `json:"data"`
	}{
		Data: resp,
	}
	c.JSON(http.StatusOK, respData)
}

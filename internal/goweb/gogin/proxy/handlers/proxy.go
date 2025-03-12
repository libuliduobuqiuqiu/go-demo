package handlers

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetProxyPassUrl(rawURL string) (proxyPassUrl *url.URL, err error) {
	tmpUrl, err := url.Parse(rawURL)
	if err != nil {
		return
	}

	queryParams := tmpUrl.Query()
	proxyPassUrlStr := queryParams.Get("proxy_pass")
	proxyPassUrl, err = url.Parse(proxyPassUrlStr)
	if err != nil {
		return nil, err
	}

	return
}

func Proxy(ctx *gin.Context) {
	w := ctx.Writer
	req := ctx.Request

	proxyPassUrl, err := GetProxyPassUrl(req.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(proxyPassUrl)
	proxy.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	proxy.Director = func(r *http.Request) {
		r.Host = proxyPassUrl.Host
		r.URL.Path = proxyPassUrl.Path
		r.URL.Scheme = proxyPassUrl.Scheme
		r.URL.Host = proxyPassUrl.Host
		r.URL.RawQuery = proxyPassUrl.RawQuery
	}

	proxy.ServeHTTP(w, req)
}

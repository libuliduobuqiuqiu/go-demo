package httpdemo

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"time"
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

func ReverseProxy(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.String())
	proxyPassUrl, err := GetProxyPassUrl(req.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(proxyPassUrl)
	proxy := httputil.NewSingleHostReverseProxy(proxyPassUrl)
	proxy.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	proxy.ServeHTTP(w, req)
}

func StartReverseProxy() {
	http.HandleFunc("/netac", ReverseProxy)
	http.ListenAndServe(":8090", nil)
}

func CommitDeviceReq(reqUrl string) (err error) {
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

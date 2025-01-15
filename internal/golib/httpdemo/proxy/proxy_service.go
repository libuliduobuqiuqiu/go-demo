package proxy

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
	proxyPassUrl, err := GetProxyPassUrl(req.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(proxyPassUrl)
	proxy := httputil.NewSingleHostReverseProxy(proxyPassUrl)
	proxy.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	proxy.Director = func(r *http.Request) {
		r.Host = proxyPassUrl.Host
		r.URL.Path = proxyPassUrl.Path
		r.URL.Scheme = proxyPassUrl.Scheme
		r.URL.Host = proxyPassUrl.Host
		r.URL.RawQuery = ""
	}

	proxy.ServeHTTP(w, req)
}

func StartReverseProxy() {
	fmt.Println("Server started on :8090")
	http.HandleFunc("/netac/base/proxy", ReverseProxy)
	http.HandleFunc("/netac/base/ssh", HandleSshConnection)
	http.ListenAndServe(":8090", nil)
}

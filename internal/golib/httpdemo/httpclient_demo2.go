package httpdemo

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func RequestDeviceAPI(reqURL, username, password string) {
	data := url.Values{}
	data.Set("file-name", "test.xml")
	data.Set("file-path", "configs/")
	reqData := data.Encode()

	fmt.Println(reqData)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, reqURL, strings.NewReader(reqData))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respData))
}

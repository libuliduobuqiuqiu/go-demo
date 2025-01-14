package proxy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
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

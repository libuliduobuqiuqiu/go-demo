package service

import (
	"io"
	"os/exec"
	"strings"
)

type NetworkService struct{}

func NewNetworkService() *NetworkService {
	return &NetworkService{}
}

func (n *NetworkService) Ping(address, timeout, count string) (res string, err error) {
	cmd := exec.Command("ping", address, "-c", count, "-W", timeout)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	defer out.Close()
	if err = cmd.Start(); err != nil {
		return
	}

	var buf = new(strings.Builder)
	if _, err = io.Copy(buf, out); err != nil {
		return
	}
	return buf.String(), nil
}

func (n *NetworkService) Traceroute(address string, count, timeout, ttl int) (res string, err error) {
	return
}

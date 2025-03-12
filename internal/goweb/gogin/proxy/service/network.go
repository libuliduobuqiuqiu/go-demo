package service

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/aeden/traceroute"
	"github.com/sirupsen/logrus"
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

func (n *NetworkService) Traceroute(address string, count, timeout, ttl int) (res []string, err error) {
	tracerouteOption := traceroute.TracerouteOptions{}
	tracerouteOption.SetRetries(count)
	tracerouteOption.SetMaxHops(ttl)
	tracerouteOption.SetTimeoutMs(timeout)

	// The mac terminal maybe occure "operation not permitted" error, which can be resolved by useing "sudo" command .
	tracerouteRes, err := traceroute.Traceroute(address, &tracerouteOption)
	if err != nil {
		logrus.Error(err)
		return
	}

	for _, hop := range tracerouteRes.Hops {
		tmp := fmt.Sprintf("%-3d %v (%v)  %v\n", hop.TTL, hop.HostOrAddressString(), hop.AddressString(), hop.ElapsedTime)
		fmt.Println(tmp)
		res = append(res, tmp)
	}

	return
}

package device

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/ssh"
)

type DeviceClient struct {
	client  *ssh.Client
	sess    *ssh.Session
	inPipe  io.WriteCloser
	outPipe io.Reader
}

func NewDeviceClient(username, password, addr string, port int) (devClient *DeviceClient, err error) {
	config := &ssh.ClientConfig{}
	config.SetDefaults()
	config.Ciphers = append(config.Ciphers, "3des-cbc")
	config.KeyExchanges = append(config.KeyExchanges, "diffie-hellman-group1-sha1")
	config.User = username
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	config.Auth = append(config.Auth, ssh.Password(password))

	c, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", addr, port), config)
	if err != nil {
		return
	}

	devClient = &DeviceClient{}
	devClient.client = c
	return
}

func (d *DeviceClient) Close() error {
	err := d.client.Close()

	if d.sess != nil {
		err = d.sess.Close()
		return err
	}

	return err
}

func (d *DeviceClient) ShellSession() error {
	sess, err := d.client.NewSession()
	if err != nil {
		return err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := sess.RequestPty("xterm", 768, 1024, modes); err != nil {
		return err
	}

	in, err := sess.StdinPipe()
	if err != nil {
		return err
	}

	out, err := sess.StdoutPipe()
	if err != nil {
		return err
	}

	err = sess.Shell()
	if err != nil {
		return err
	}

	d.inPipe = in
	d.outPipe = out

	d.ReadLine("")
	return nil
}

func (d *DeviceClient) SendCmd(cmd string, extendCmd ...string) (res string, err error) {

	d.inPipe.Write([]byte(cmd))
	d.ReadLine(cmd, extendCmd...)

	return
}

func (d *DeviceClient) ReadLine(cmd string, extendCmd ...string) (res string, err error) {
	var content strings.Builder
	buf := make([]byte, 2048)
	for {

		n, err := d.outPipe.Read(buf)
		if err != nil {
			return "", err
		}

		b := string(buf[:n])
		content.WriteString(b)
		fmt.Println(n, b)

		if len(extendCmd) > 0 {
			d.inPipe.Write([]byte(extendCmd[0]))
		}

		if strings.Contains(b, ">") || strings.HasSuffix(b, "#") {
			break
		}
	}

	res = content.String()
	return
}

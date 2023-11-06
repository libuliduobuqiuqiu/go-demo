package main

import (
	"fmt"
	"sunrun/ConcurrencyDemo"
	"sync"
)

type ExecInterface interface {
	Hello() string
	Speak() string
}

type Handler interface {
	NewExecTask() ExecInterface
}

type Device struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func (d *Device) Hello() string {
	return d.Name
}

func (d *Device) Speak() string {
	return d.Address
}

func resetTaskStatus(d *Device, err *error) {

	fmt.Println("Hello,world", err)
	if err != nil {
		fmt.Println(*err)
		fmt.Println(d.Hello())
		fmt.Println(d.Address)
	}

}

func genError() (string, error) {
	return "", nil
}

func execTask() {

	var err error
	var h string
	d := Device{"127.0.0.1", "Server"}

	defer resetTaskStatus(&d, &err)

	if h, err = genError(); err != nil {
		d.Name = "Error Server"
		fmt.Println(h)
		fmt.Println(d.Name)
		return
	}

}

type PersonInfo struct {
	Info map[string]string
}

var person PersonInfo

func write(p PersonInfo, key string, value string) {
	p.Info[key] = value
}

func read(p PersonInfo, key string) {
	if v, ok := p.Info[key]; ok {
		fmt.Println(key, ":", v)
	}
}

func write_person(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("h%d", i)
		value := fmt.Sprintf("v%d", i)
		write(person, key, value)
		read(person, key)
	}
}

func print_person(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		if v, ok := person.Info[fmt.Sprintf("h%d", i)]; ok {
			fmt.Println(v)
		}
	}
}

func main() {
	ConcurrencyDemo.DeviceExecCommands()
}

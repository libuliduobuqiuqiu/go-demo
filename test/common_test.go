package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"godemo/internal/godemo"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"
)

var (
	Param string
)

// Test the slice expansion mechanism.
func TestSliceAppend(t *testing.T) {
	var b []int = nil
	fmt.Printf("%p\n", b)
	fmt.Println(len(b), cap(b))

	b = append(b, 100)
	b = append(b, 200, 200, 200, 200, 200, 200)
	fmt.Printf("%p\n", b)
	fmt.Println(len(b), cap(b))

	s := make([]int, 2, 2)
	s[0] = 1
	s[1] = 2

	// 打印初始切片的地址和底层数组的地址
	fmt.Printf("Initial slice address: %p\n", s)
	fmt.Printf("Initial array address: %p\n", unsafe.Pointer(&s[0]))

	// 添加元素，触发扩容
	s = append(s, 3)

	// 打印扩容后的切片地址和底层数组的地址
	fmt.Printf("New slice address: %p\n", s)
	fmt.Printf("New array address: %p\n", unsafe.Pointer(&s[0]))
}

// Test whether the loop has fixed the shared variable issue.
func TestRange(t *testing.T) {
	var (
		a     []int
		funcs []func()
	)

	a = []int{1, 2, 3, 4, 5, 6}

	for _, v := range a {
		fmt.Printf("v 的地址：%p\n", &v)
		funcs = append(funcs, func() {
			fmt.Println(v)
		})
	}

	for _, f := range funcs {
		f()
	}

	fmt.Println("--------------------")

	sg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		sg.Add(1)
		go func() {
			defer sg.Done()
			fmt.Println(i)
		}()
	}

	sg.Wait()
}

func TestComb(t *testing.T) {

	var k = 3
	var nums = []int{1, 2, 3, 4, 5}

	var res ([][]int)
	var breakTrack func(start int, comb []int)

	breakTrack = func(start int, comb []int) {
		if len(comb) == k {
			res = append(res, comb)
		}

		for i := start; i < len(nums); i++ {
			breakTrack(i+1, append(comb, nums[i]))
		}
	}

	breakTrack(0, []int{})

	for _, r := range res {
		fmt.Println(r)
	}
}

func TestDeleteSlice(t *testing.T) {

	a := []string{"hello,", "world", "ifocnifg"}
	fmt.Println(len(a))
	a = a[1:]

	fmt.Println(len(a))

	b := make(map[string][]string)
	b["info"] = append(b["info"], "name")
	fmt.Println(b["info"])

}

func TestUnmarshal(t *testing.T) {

	text := "[{\"type\":\"contains\",\"value\":\"\n\",\"key\":\"Password:\"},{\"type\":\"contains\",\"value\":\"\",\"key\":\"Horizon-86(LICENSE-EXPIRES IN 81 DAYS)#\"}]"

	text = strings.ReplaceAll(text, "\"\n\"", "\"\\n\"")
	fmt.Println(text)

	quote := strconv.Quote(text)
	fmt.Println(quote)

	type option struct {
		Type  string `json:"type"`
		Value string `json:"value"`
		Key   string `json:"key"`
	}

	var o []option
	if err := json.Unmarshal([]byte(text), &o); err != nil {
		t.Fatal(err)
	}

	for _, t := range o {
		fmt.Println(t)
	}

}

func switchType(p any) {
	switch v := p.(type) {
	case string:
		fmt.Println(v)
	case []byte:
		fmt.Println(string(v))

	}

}

func TestSwitchType(t *testing.T) {

	a := "hello,world"

	switchType(a)
	switchType([]byte(a))

}

func TestLazyError(t *testing.T) {
	godemo.LazyGetError()
}

func TestCountTime(t *testing.T) {

	err := filepath.Walk("/data/Company/log/netac/oarsflow/", func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		subTime := time.Now().Sub(info.ModTime())
		subDays := subTime.Hours() / 24
		if subDays > 30 {
			fmt.Println(filePath, info.Name())
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

}

func ReturnParam() (p string, err error) {
	return "", errors.New("test")
}

func TestUseParam(t *testing.T) {
	var (
		Param string
	)
	fmt.Println(&Param)

	Param, err := ReturnParam()
	fmt.Println(err)
	fmt.Println(&Param)

}

func TestUseNet(t *testing.T) {

	_, ipnet, err := net.ParseCIDR("192.168.5.0/24")
	if err != nil {
		t.Fatal()
	}

	networkIP := ipnet.IP.Mask(ipnet.Mask)
	inc(networkIP)

	for ip := networkIP; ipnet.Contains(ip); inc(ip) {
		fmt.Println(ip.String())
	}

}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func TestRegxp(t *testing.T) {
	reg := regexp.MustCompile(`\(serverIpAddr=([\w\.]+)&&serverPort=([\w\.]+)\)`)
	tmpStr := `(serverIpAddr=192.168.1.210&&serverPort=1699)||(serverIpAddr=192.168.1.237&&serverPort=11499)`

	res := reg.FindAllStringSubmatch(tmpStr, -1)
	fmt.Println(res)

}

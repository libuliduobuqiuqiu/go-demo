package godemo

import (
	"bufio"
	"embed"
	"fmt"
	"log"
)

//go:embed time_demo.go
var f embed.FS

func GetEmbedFile() {
	file, err := f.Open("time_demo.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	buf := bufio.NewScanner(file)

	for buf.Scan() {
		fmt.Println(buf.Text())
	}
}

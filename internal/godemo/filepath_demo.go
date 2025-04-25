package godemo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetAbsPath() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rootPath)
	tmp, err := filepath.Abs("../../internal/dao")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmp)

}

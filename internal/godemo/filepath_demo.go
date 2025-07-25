package godemo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	ModifiedTime time.Time `json:"modified_time"`
}

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

func ScanDir(dirPath string) ([]*FileInfo, error) {
	var files []*FileInfo
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		files = append(files, &FileInfo{
			Name:         info.Name(),
			Size:         info.Size(),
			ModifiedTime: info.ModTime(),
		})

		return nil
	})

	return files, err
}

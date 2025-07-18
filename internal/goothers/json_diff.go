package goothers

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/google/go-cmp/cmp"
)

func DiffJson(localFile, remoteFile string) error {
	var (
		localObj  any
		remoteObj any
	)

	localData, err := os.ReadFile(localFile)
	if err != nil {
		return err
	}

	remoteData, err := os.ReadFile(remoteFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(localData, &localObj); err != nil {
		return err
	}

	if err = json.Unmarshal(remoteData, &remoteObj); err != nil {
		return err
	}

	localFileName := path.Base(localFile)
	remoteFileName := path.Base(remoteFile)

	if diff := cmp.Diff(localObj, remoteObj); diff != "" {
		fmt.Println("差异:\n", diff)
	} else {
		fmt.Printf("%s 和 %s 完全相等\n", localFileName, remoteFileName)
	}
	return nil
}

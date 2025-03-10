package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

func PrettyJson(data []byte) {
	var raw interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		logrus.Warnf("Unmarshal %s error: %v", string(data), err)
		return
	}

	formatted, err := json.MarshalIndent(raw, "", "   ")
	if err != nil {
		logrus.Warnf("Marshal json error: %v", err)
	}

	fmt.Println(string(formatted))
}

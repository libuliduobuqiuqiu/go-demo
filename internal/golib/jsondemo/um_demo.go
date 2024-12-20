package jsondemo

import (
	"encoding/json"
	"fmt"
)

type obj struct {
	Num1 int                    `json:"num1"`
	Data map[string]interface{} `json:"data"`
}

func UnmarshalObj() error {

	body := `{
    "num1": 11,
    "Data":{
            "max_length": 12
        }
	}`
	var tmp obj
	if err := json.Unmarshal([]byte(body), &tmp); err != nil {
		return err
	}

	fmt.Printf("%T, %v\n", tmp.Num1, tmp.Num1)
	fmt.Printf("%T, %v\n", tmp.Data["max_length"], tmp.Data["max_length"])
	return nil
}


package goothers

import (
	"encoding/json"
	"os"
	"regexp"

	"github.com/bwmarrin/snowflake"
)

var (
	uuidRegex = regexp.MustCompile("^[a-fA-F0-9-]{32}$")
	mapping   = make(map[string]string)
)

func ReplaceUUIDWithSnowflakeID(fileName string) (err error) {

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	beforeData, err := BeforeExec(fileData)
	if err != nil {
		return
	}

	var data any
	if err = json.Unmarshal(beforeData, &data); err != nil {
		return
	}

	node, err := snowflake.NewNode(67)
	if err != nil {
		return
	}

	CollectUUID(data, mapping, node)

	ReplaceUUID(data, mapping)

	out, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = os.WriteFile(fileName, out, 0755)
	return
}

func BeforeExec(fileData []byte) (res []byte, err error) {

	var flowConfig FlowConfig

	if err = json.Unmarshal(fileData, &flowConfig); err != nil {
		return
	}

	for _, node := range flowConfig.FlowChainNodes {
		if node.ID != node.NodeTag {
			node.ID = node.NodeTag
		}
	}

	res, err = json.Marshal(flowConfig)
	return
}

func CollectUUID(data any, mapping map[string]string, node *snowflake.Node) {
	switch vv := data.(type) {
	case map[string]any:
		for _, v := range vv {
			CollectUUID(v, mapping, node)
		}
	case []any:
		for _, v := range vv {
			CollectUUID(v, mapping, node)
		}
	case string:
		if uuidRegex.MatchString(vv) {
			if _, ok := mapping[vv]; !ok {
				mapping[vv] = node.Generate().String()
			}
		}
	}
}

func ReplaceUUID(data any, mapping map[string]string) {
	switch vv := data.(type) {
	case map[string]any:
		for k, v := range vv {
			if s, ok := v.(string); ok && mapping[s] != "" {
				vv[k] = mapping[s]
			} else {
				ReplaceUUID(v, mapping)
			}
		}
	case []any:
		for i, v := range vv {
			if s, ok := v.(string); ok && mapping[s] != "" {
				vv[i] = mapping[s]
			} else {
				ReplaceUUID(v, mapping)
			}
		}
	}
}

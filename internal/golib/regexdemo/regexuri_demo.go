package regexdemo

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// ${http_type}${address}/adcapi/v2.0?authkey=%24%7Bkey%7D&action=logout
func ReplaceUri(uri string) error {
	var err error

	uri, err = url.QueryUnescape(uri)
	if err != nil {
		return err
	}

	fmt.Println(uri)

	reg := regexp.MustCompile(`\$\{\w+\}`)
	varNames := reg.FindAllStringSubmatch(uri, -1)

	var keys []string
	for _, vars := range varNames {
		if len(vars) > 0 {
			keys = append(keys, vars[0])
		}
	}

	urlMap := map[string]string{
		"${http_type}": "http://",
		"${address}":   "127.0.0.1",
		"${key}":       "admin",
	}

	for _, key := range keys {
		if v, ok := urlMap[key]; ok {
			uri = strings.ReplaceAll(uri, key, v)
		}
	}

	fmt.Println(uri)

	return nil
}

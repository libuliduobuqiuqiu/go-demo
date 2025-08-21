package golib

import (
	"fmt"
	"godemo/internal/golib/regexdemo"
	"strings"
	"testing"
)

func TestRgexpReplace(t *testing.T) {

	// jsContent := `const regex = new RegExp("lb node .*?\\n\\w", 'gms')`
	// replaceStr := regexdemo.PatchDotAllRegexp(jsContent)
	// fmt.Println(replaceStr)
	//
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic test",
			input:    `const regex = new RegExp("lb node .*?\\n\\w", 'gms')`,
			expected: `const regex = new RegExp("lb node [\s\S]*?\\n\\w",'gm')`,
		},
		{
			name:     "basic test2",
			input:    `const regex = new RegExp('lb node .*?\\n\\w', "gms")`,
			expected: `const regex = new RegExp('lb node [\s\S]*?\\n\\w',"gm")`,
		},
		{
			name:     "nomal regexp",
			input:    `const regex = new RegExp("lb node \\n\\w", 'gm')`,
			expected: `const regex = new RegExp("lb node \\n\\w", 'gm')`,
		}, {
			name:     "logn regexp",
			input:    `tmpPort = Js.getGroup(new RegExp("lb node .*?(port.*?)\\n\\w", 'gms'), node)`,
			expected: `tmpPort = Js.getGroup(new RegExp("lb node [\s\S]*?(port[\s\S]*?)\\n\\w",'gm'), node)`,
		},
		{
			name: "multiple regex",
			input: `name = Js.getGroup(new RegExp("lb node (.*?) ", 'gms'), tmpNode) //节点名称
            address = Js.getGroup(new RegExp("lb node .*?\\s(.*?)\\s", 'gms'), tmpNode) //节点ip`,
			expected: `name = Js.getGroup(new RegExp("lb node ([\s\S]*?) ",'gm'), tmpNode) //节点名称
            address = Js.getGroup(new RegExp("lb node [\s\S]*?\\s([\s\S]*?)\\s",'gm'), tmpNode) //节点ip`,
		}, {
			name:     "reliteral regex",
			input:    `let commonNameRegex = /CN\s*=\s*([^,\n]+)/gm;`,
			expected: `let commonNameRegex = /CN\s*=\s*([^,\n]+)/gm;`,
		}, {
			name:     "reliteral regex 2",
			input:    `let serialNumberRegex = /Serial Number:\s*(\S+)/gms;`,
			expected: `let serialNumberRegex = /Serial Number:\s*(\S+)/gm;`,
		}, {
			name:     "reliteral regex 3",
			input:    `let serialNumberRegex = /Serial Number:\s.*(\S+)/gms;`,
			expected: `let serialNumberRegex = /Serial Number:\s[\s\S]*(\S+)/gm;`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := regexdemo.PatchDotAllRegexp(tt.input)

			if got != tt.expected {
				fmt.Println(got)
				fmt.Println(tt.expected)
				t.Errorf("%s PatchDotAllRegexp() = %q want %q", tt.name, got, tt.expected)
			}
		})
	}

}

func TestReplaceUri(t *testing.T) {

	uri := "${http_type}${address}/adcapi/v2.0?authkey=%24%7Bkey%7D&action=logout"

	err := regexdemo.ReplaceUri(uri)
	if err != nil {
		t.Fatal(err)
	}

}

func TestPattern(t *testing.T) {
	a := `"abc.conf"`
	a = strings.ReplaceAll(a, ".", `[\s\S]`)
	fmt.Println(a)

}

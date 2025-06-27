package regexdemo

import (
	"fmt"
	"regexp"
	"strings"
)

func PatchDotAllRegexp(js string) string {
	reg := regexp.MustCompile(`new\s*RegExp\(\s*(['"].*?['"])\s*,\s*(['"].*?s.*?['"])\s*\)`)
	js = reg.ReplaceAllStringFunc(js, func(match string) string {
		sub := reg.FindStringSubmatch(match)
		if len(sub) != 3 {
			return match
		}

		pattern := sub[1]
		flags := sub[2]
		newFlags := strings.ReplaceAll(flags, "s", "")
		newPattern := strings.ReplaceAll(pattern, ".", "[\\s\\S]")
		return `new RegExp(` + newPattern + `,` + newFlags + `)`
	})

	reLiteral := regexp.MustCompile(`([=(:\s])/(.*?)/([gimsuy]*)`)
	js = reLiteral.ReplaceAllStringFunc(js, func(match string) string {
		sub := reLiteral.FindStringSubmatch(match)
		if len(sub) != 4 {
			return match
		}
		prefix := sub[1]
		pattern := sub[2]
		flags := sub[3]

		if !strings.Contains(flags, "s") {
			return match
		}

		newFlags := strings.ReplaceAll(flags, "s", "")
		newPattern := strings.ReplaceAll(pattern, ".", "[\\s\\S]")

		return fmt.Sprintf(`%s/%s/%s`, prefix, newPattern, newFlags)
	})

	return js
}

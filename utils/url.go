package utils

import "regexp"

var r = regexp.MustCompile(`https?://[\w/:%#\$&\?\(\)~\.=\+\-]+`)

func IsUrl(target string) bool {
	if len(target) == 0 {
		return false
	}

	return r.Match([]byte(target))
}

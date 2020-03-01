package 剑指Offer2

import "strings"

func replaceSpace(s string) string {
	if s == "" {
		return s
	}

	s = strings.ReplaceAll(s, " ", "%20")
	return s
}

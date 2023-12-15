package util

import "strings"

func JoinKey(keys ...string) string {
	return strings.Join(keys, ":")
}

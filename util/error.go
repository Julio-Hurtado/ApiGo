package util

import "strings"

func FormatearError(err string) string {
	mensaje := strings.Split(err, ":")[1]
	return mensaje
}

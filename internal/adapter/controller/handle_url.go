package controller

import (
	"fmt"
	"strings"
)

func HandleUrl(prefix, url string) string {
	path := strings.TrimPrefix(url, prefix)
	path = strings.TrimSpace(path)

	fmt.Println(prefix, url)

	if strings.Contains(path, "?") {
		return "?"
	}

	return path
}

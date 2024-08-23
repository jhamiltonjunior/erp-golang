package controller

import "strings"

func HandleUrl(prefix, url string) string {
	path := strings.TrimPrefix(url, prefix)
	path = strings.TrimSpace(path)

	if strings.Contains(path, "?") {
		pathSlice := strings.Split(path, "?")
		return pathSlice[0]
	}

	return path
}

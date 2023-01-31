package folderstructure

import "regexp"

func FileSystemEscape(input string) string {
	pattern, _ := regexp.Compile(`[^\w\d.+-]`)
	return pattern.ReplaceAllString(input, "_")
}

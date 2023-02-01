package folderstructure

import "regexp"

// FileSystemEscape removes all illegal characters from the input string so that it's safe to use for file system operations
func FileSystemEscape(input string) string {
	pattern, _ := regexp.Compile(`[^\w\d.+-]`)
	return pattern.ReplaceAllString(input, "_")
}

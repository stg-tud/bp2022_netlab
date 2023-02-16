package folderstructure

import (
	"os"

	logger "github.com/gookit/slog"
)

// OverwriteExisting determines whether existing files should be overwritten (true) or left untouched (false)
var OverwriteExisting bool = false

// MayCreateFile returns whether a file may be created/overwritten or should not be touched.
func MayCreatePath(path string) bool {
	logger.Tracef("Checking writing permissions for path \"%s\"", path)
	if OverwriteExisting {
		return true
	}
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

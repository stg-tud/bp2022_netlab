package folderstructure

import (
	"os"

	logger "github.com/gookit/slog"
)

// The name of the environment variable to check if files may be overwritten
const SkipExistingEnv string = "SKIP_EXISTING"

// MayCreateFile returns whether a file may be created/overwritten or should not be touched.
func MayCreatePath(path string) bool {
	logger.Tracef("Checking writing permissions for path \"%s\"", path)
	if os.Getenv(SkipExistingEnv) != "1" {
		return true
	}
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

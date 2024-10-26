package projectx

import (
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

// GetGoModPath returns the deepest directory containing the go.mod file
// in the working directory parents
func GetGoModPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "os.Getwd failed")
	}
	goModExists := false

	for !goModExists {
		if wd == "/" {
			return "", errors.New("go.mod not found")
		}
		goModExists, err = IsExist(path.Join(wd, "go.mod"))
		if err != nil {
			return "", errors.Wrap(err, "cannot check go.mod existence")
		}
		if goModExists {
			break
		}
		wd = filepath.Dir(wd)
	}

	// reached project root
	return wd, nil
}

// IsExist checks if the given file or directory exists
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errors.Wrap(err, "os.Stat failed")
}

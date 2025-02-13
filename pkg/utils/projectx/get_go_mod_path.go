package projectx

import (
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

// GetGoModPath returns the deepest directory containing the go.mod file
// in the working directory or its parents. It starts from the current
// working directory and traverses up until it finds a go.mod file or
// reaches the root directory.
//
// Returns an error if:
// - Cannot get current working directory
// - Cannot check file existence
// - No go.mod file found in any parent directory
//
// Example:
//
//	// If current directory is /home/user/projects/myapp/cmd
//	// and go.mod is in /home/user/projects/myapp
//	path, err := GetGoModPath()
//	// path will be "/home/user/projects/myapp"
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

// IsExist checks if the given file or directory exists.
// Returns true if the path exists, false if it doesn't exist,
// and an error if the check fails for other reasons (e.g., permissions).
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

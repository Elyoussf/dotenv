package dotenvlocator

import (
	"fmt"
	"os"
	"path/filepath"
)

func LocateDotEnv(startDir, filename string) (string, error) {
	curr := startDir
	for {
		fullpath := filepath.Join(curr, filename)

		_, err := os.Stat(fullpath)

		if err == nil {
			return fullpath, nil
		}
		if !os.IsNotExist(err) {
			return "", err // maybe permission error or something at this case .env could never be found!
		}
		parent := filepath.Dir(curr)

		if parent == curr {
			// reaching a point where there is no upward go (e.g reaching the root on Linux )
			return "", fmt.Errorf("dotenv could not be found")
		}
		curr = parent
	}
}

package word

import (
	"os"
	"path/filepath"
)

func Root() string {
	current, _ := os.Getwd()
	for {
		goModPath := filepath.Join(current, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return current // Found go.mod, so this is the root
		}

		// Move up one directory
		parent := filepath.Dir(current)
		if parent == current { // Reached root of filesystem
			return current
		}
		current = parent
	}
}

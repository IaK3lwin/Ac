package storage

import (
	"os"
	"path/filepath"
	"strings"
)

func Expand(path string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	switch {
	case path == "~":
		return home

	case strings.HasPrefix(path, "~/"):
		return filepath.Join(home, path[2:])
	}

	return path
}

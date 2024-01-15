package utils

import (
	"os"
	"path/filepath"
)

func MakeBasePath() (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cur, "..", "..", "..", "..", "persons"), nil
}

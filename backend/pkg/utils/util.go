package utils

import (
	"log"
	"os"
	"path/filepath"
)

func MakeBasePath() string {
	cur, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(cur, "..", "..", "..", "..", "persons")
}

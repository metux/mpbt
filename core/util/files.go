package util

import (
	"os"
)

func FileExists(fn string) bool {
	_, err := os.Stat(fn)
	return err == nil
}

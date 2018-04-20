package pathfinder

import (
	"os"
)

func DoesExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

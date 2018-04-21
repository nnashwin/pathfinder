package pathfinder

import (
	"log"
	"os"
	"path/filepath"
)

func DoesExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreatePath(path string) error {
	if DoesExist(path) == true {
		log.Fatal("You can not create a path that already exists.")
	}

	dir, file := filepath.Split(path)
	if file == "" {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	} else {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		newFile, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()

		return nil
	}
}

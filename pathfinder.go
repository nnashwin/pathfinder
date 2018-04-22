package pathfinder

import (
	"errors"
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

func CreateFile(path string) error {
	if DoesExist(path) == true {
		return errors.New("You can not create a file that already exists.")
	}

	dir, _ := filepath.Split(path)
	if dir != "" {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer newFile.Close()

	return nil
}

func CreateDir(path string) error {
	if DoesExist(path) == true {
		return errors.New("You can not create a directory that already exists.")
	}

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil

}

func GetDir(path string) (dir string) {
	dir, _ = filepath.Split(path)

	return dir
}

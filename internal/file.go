package file

import (
	"os"
)

func Create(name string) (*os.File, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteFile(name string, data []byte, perm os.FileMode) error {
	err := os.WriteFile(name, data, perm)
	if err != nil {
		return err
	}
	return nil
}

func Delete(name string) error {
	err := os.Remove(name)
	if err != nil {
		return err
	}
	return nil
}

func MakeDirectory(dirName string, perm os.FileMode) error {
	err := os.Mkdir(dirName, perm)
	if err != nil {
		return err
	}
	return nil
}

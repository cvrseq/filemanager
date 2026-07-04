package file

import (
	"os"
	"time"
)

type Object struct {
	file *File
	dir  *Directory
}

type Directory struct {
	name      string
	writes    string
	filenames string
}

type Metadata struct {
	filename string
	size     uint64
	kind     *Object
}

type File struct {
	name        string
	size        uint64
	permissions string
	owner       string
	created_at  time.Time
	metadata    *Metadata
}

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

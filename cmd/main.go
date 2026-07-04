package main

import (
	"os"

	file "github.com/cvrseq/filemanager/internal"
)

func main() {

	var (
		dirName string
		name    string
		data    []byte
		perm    os.FileMode
	)

	dirName = "test"
	name = "test_file.txt "
	data = []byte("hello world")
	perm = os.FileMode(0644)

	file.Create(name)
	file.WriteFile(name, data, perm)
	file.MakeDirectory(dirName, perm)
	file.Delete(name)
	file.Delete(dirName)
}

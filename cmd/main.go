package main

import (
	"os"

	file "github.com/cvrseq/filemanager/internal"
)

func main() {
	var (
		name string
		data []byte
		perm os.FileMode
	)

	name = "test_file"
	data = []byte("hello world")
	perm = os.FileMode(0644)

	file.Create(name)
	file.WriteFile(name, data, perm)

}

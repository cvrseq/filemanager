package main

import (
	file "github.com/cvrseq/filemanager/internal"
)

func main() {
	var name string
	name = "test_file"

	file.Create(name)

}

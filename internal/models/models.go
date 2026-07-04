package models

import "time"

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

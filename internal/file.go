package file

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type FileInfo struct { 
	Name string  `json:"name"`
	Permission string `json:"perm"`
	Data string `json:"data"`
	Type string `json:"type"`
	Size uint64 `json:"size"`
}

type DirRequest struct {
	Name string      `json:"name"`
	Perm os.FileMode `json:"permission"`
}

type DirResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type FileRequest struct {
	Name string `json:"name"`
}

type FileResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type DeleteRequest struct {
	Name string `json:"name"`
}

type DeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CreateAndWriteRequest struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
	Perm os.FileMode `json:"permission"`
}

type CreateAndWriteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


func CreateFileHandler(w http.ResponseWriter, r *http.Request) {
	var req FileRequest
	decodeHelperByCreateFile(&req, r)

	_, err := os.Create(req.Name)
	if err != nil {
		log.Printf("Create failed: %v", err)
	}

	encodeHelperByCreateFile(w, r)
}

func CreateDirHandler(w http.ResponseWriter, r *http.Request) {
	var req DirRequest
	decodeHelperByMakeDirectory(&req, r)

	err := os.Mkdir(req.Name, req.Perm)
	if err != nil {
		log.Printf("Make directory failed: %v", err)
	}

	encodeHelperByMakeDirectory(w, r)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	decodeHelperByDeleteFileOrDir(&req, r)

	err := os.Remove(req.Name)
	if err != nil {
		log.Printf("Delete file or directory failed: %v", err)
	}
	encodeHelperByDeleteFileOrDir(w, r)
}

func CreateAndWriteFileIfExistHandler(w http.ResponseWriter, r *http.Request) { 
	var req CreateAndWriteRequest
	decodeHelperByCreateAndWrite(&req, r) 

	err := os.WriteFile(req.Name, req.Data, req.Perm)
	if err != nil {
		log.Printf("Write file failed: %v", err)
	}
	encodeHelperByCreateAndWrite(w, r)
}



func decodeHelperByMakeDirectory(req *DirRequest, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}
	return nil
}

func encodeHelperByMakeDirectory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(DirResponse{
		Status:  "created",
		Message: "Make directory handler successed",
	})
}

func decodeHelperByCreateFile(req *FileRequest, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}
	return nil
}

func encodeHelperByCreateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(FileResponse{
		Status:  "created",
		Message: "Create file handler successed",
	})
}

func decodeHelperByDeleteFileOrDir(req *DeleteRequest, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}
	return nil
}

func encodeHelperByDeleteFileOrDir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	json.NewEncoder(w).Encode(DeleteResponse{
		Status:  "Deleted",
		Message: "Delete file handler successed",
	})
}

func decodeHelperByCreateAndWrite(req *CreateAndWriteRequest, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}
	return nil
}

func encodeHelperByCreateAndWrite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	json.NewEncoder(w).Encode(CreateAndWriteResponse{
		Status:  "Created and writed",
		Message: "Create and write file handler successed",
	})
}
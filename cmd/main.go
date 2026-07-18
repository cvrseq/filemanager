package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	file "github.com/cvrseq/filemanager/internal"
)

func main() {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	mux.HandleFunc("POST /create_file", file.CreateFileHandler)
	mux.HandleFunc("POST /create_directory", file.CreateDirHandler)
	mux.HandleFunc("POST /write_file", file.CreateAndWriteFileIfExistHandler)
	mux.HandleFunc("DELETE /delete", file.DeleteHandler)

	go func() {
		log.Printf("Server starting on %s", srv.Addr)
		if err := http.ListenAndServe(srv.Addr, mux); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server ... ")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server stopped gracefully")
}

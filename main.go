package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed static
	content embed.FS
)

func fileServer() http.Handler {
	dir, err := fs.Sub(content, "static")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(dir))
}

func main() {
	http.Handle("/", fileServer())
	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

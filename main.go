package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const filePathRoot = "." // root directory

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot))))
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Server listening from %s on port: %s\n", filePathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

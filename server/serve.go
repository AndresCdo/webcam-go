package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

const dir = "."

func main() {
	// Get the port from the environment variable, defaulting to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize the file server
	fs := http.FileServer(http.Dir(dir))

	// Log the server start-up message
	log.Printf("Serving %s on http://localhost:%s\n", dir, port)

	// Define the custom HTTP handler
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("Content-Type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	})

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

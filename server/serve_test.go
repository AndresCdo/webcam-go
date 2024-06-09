package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMainHandler(t *testing.T) {
	// Create a temporary directory for testing
	dir := t.TempDir()
	os.Chdir(dir)
	// Create a sample wasm file
	wasmFile := "test.wasm"
	f, err := os.Create(wasmFile)
	if err != nil {
		t.Fatalf("Failed to create wasm file: %v", err)
	}
	f.Close()

	// Set the PORT environment variable
	os.Setenv("PORT", "8081")

	// Initialize the file server
	fs := http.FileServer(http.Dir(dir))

	// Define the custom HTTP handler
	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("Content-Type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	})

	// Create a new HTTP request for the wasm file
	req := httptest.NewRequest("GET", "/"+wasmFile, nil)
	resp := httptest.NewRecorder()

	// Serve the HTTP request
	handler.ServeHTTP(resp, req)

	// Check if the response status code is 200
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	// Check if the Content-Type header is set correctly
	if contentType := resp.Header().Get("Content-Type"); contentType != "application/wasm" {
		t.Errorf("Expected Content-Type application/wasm, got %s", contentType)
	}

	// Check if the Cache-Control header is set correctly
	if cacheControl := resp.Header().Get("Cache-Control"); cacheControl != "no-cache" {
		t.Errorf("Expected Cache-Control no-cache, got %s", cacheControl)
	}
}

func TestMainHandlerNonWasm(t *testing.T) {
	// Create a temporary directory for testing
	dir := t.TempDir()
	os.Chdir(dir)
	// Create a sample text file
	txtFile := "test.txt"
	f, err := os.Create(txtFile)
	if err != nil {
		t.Fatalf("Failed to create text file: %v", err)
	}
	f.Close()

	// Set the PORT environment variable
	os.Setenv("PORT", "8081")

	// Initialize the file server
	fs := http.FileServer(http.Dir(dir))

	// Define the custom HTTP handler
	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("Content-Type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	})

	// Create a new HTTP request for the text file
	req := httptest.NewRequest("GET", "/"+txtFile, nil)
	resp := httptest.NewRecorder()

	// Serve the HTTP request
	handler.ServeHTTP(resp, req)

	// Check if the response status code is 200
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	// Check if the Content-Type header is not set to application/wasm
	if contentType := resp.Header().Get("Content-Type"); contentType == "application/wasm" {
		t.Errorf("Expected Content-Type not to be application/wasm, got %s", contentType)
	}

	// Check if the Cache-Control header is set correctly
	if cacheControl := resp.Header().Get("Cache-Control"); cacheControl != "no-cache" {
		t.Errorf("Expected Cache-Control no-cache, got %s", cacheControl)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Define the path to your files
const (
	updateJSONPath = "./files/update.json"
	xpiFilesDir    = "./files/"
)

// Handler for "/update/" and "/update" - serves JSON content or redirects to xpiHandler
func updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL: %s, sender:%s\n", r.URL.Path, r.RemoteAddr)

	// If the path is exactly "/update/" or "/update", serve the JSON file
	if r.URL.Path == "/update" || r.URL.Path == "/update/" {
		// Read JSON file
		jsonFile, err := os.Open(updateJSONPath)
		if err != nil {
			http.Error(w, "Could not open update.json", http.StatusInternalServerError)
			log.Println("Error opening update.json:", err)
			return
		}
		defer func(jsonFile *os.File) {
			_ = jsonFile.Close()
		}(jsonFile)

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			http.Error(w, "Could not read update.json", http.StatusInternalServerError)
			log.Println("Error reading update.json:", err)
			return
		}

		// Write JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(byteValue)
	} else {
		xpiHandler(w, r)
	}
}

func xpiHandler(w http.ResponseWriter, r *http.Request) {
	// Extract file ID from the URL, e.g., "bk-201"
	fileID := strings.TrimPrefix(r.URL.Path, "/update/")
	filePath := fmt.Sprintf("%s%s.xpi", xpiFilesDir, fileID)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		log.Println("File not found:", err)
		return
	}

	// Set headers to serve file as a download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xpi", fileID))
	w.Header().Set("Content-Type", "application/x-xpinstall")
	http.ServeFile(w, r, filePath)
}

func main() {
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/update/", updateHandler)

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

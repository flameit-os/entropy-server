package main

/*
FlameIT - Immersion Cooling - Entropy Server
Author: Paweł 'felixd' Wojciechowski
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	maxBytes     = 100 * 1024 * 1024 * 1024 // 100 GB
	defaultBytes = 512                      // 4096 bits
	port         = 8080
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queryBytes := r.URL.Query().Get("bytes")
		numBytes := int64(defaultBytes)

		if queryBytes != "" {
			n, err := strconv.ParseInt(queryBytes, 10, 64)
			if err == nil && n > 0 && n <= maxBytes {
				numBytes = int64(n)
			}
		}

		randomFile, err := os.Open("/dev/random")
		if err != nil {
			http.Error(w, "Failed to read from /dev/random", http.StatusInternalServerError)
			return
		}
		defer randomFile.Close()

		log.Printf("Request to provide %d bytes", numBytes)

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=random.bits")
		w.Header().Set("Content-Length", strconv.FormatInt(numBytes, 10))

		bytesWritten, err := io.CopyN(w, randomFile, int64(numBytes))
		if err != nil {
			http.Error(w, "Failed to read from /dev/random", http.StatusInternalServerError)
			return
		}

		log.Printf("Bytes provided: %d\n", bytesWritten)

	})

	log.Printf("Listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
}

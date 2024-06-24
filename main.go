package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http" // file for downloading
	"time"
)

// file for downloading
const RemoteFileURL = "https://scontent.xx.fbcdn.net/v/t39.16592-6/10000000_259873120522226_705893978463060219_n.dmg/WhatsApp-2.24.11.85.dmg?_nc_cat=103&ccb=1-7&_nc_sid=8f3826&_nc_ohc=ioEc3im-6LUQ7kNvgEI8FPg&_nc_ht=scontent.xx&oh=00_AYCP-B8PKb-5ZIK98VAxXCq2J3crJVp0J1sPZKuU28PNvA&oe=667F86CE"

// DownloadHandler handles the download speed test
func TestDownloadSpeed() {
	start := time.Now()

	response, err := http.Get(RemoteFileURL)
	if err != nil {
		fmt.Println(errors.New("error fetching remote file"))
		return
	}
	defer response.Body.Close()

	bytesWritten, err := io.Copy(io.Discard, response.Body)
	if err != nil {
		fmt.Println(errors.New("error downloading remote file"))
		return
	}
	duration := time.Since(start).Seconds()

	speed := float64(bytesWritten) / duration / 1024 / 1024 // Speed in MB/s

	//print download speed
	fmt.Printf("Download: %d bytes in %.2f seconds = %.2f MB/s\n", bytesWritten, duration, speed)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Start timing the upload

	bytesRead, err := io.Copy(io.Discard, r.Body)
	if err != nil {
		http.Error(w, "Error reading upload", http.StatusInternalServerError)
		return
	}

	duration := time.Since(start).Seconds() // Calculate how long the upload took

	speed := float64(bytesRead) / duration / 1024 / 1024 // Calculate upload speed in MB/s

	// print with upload details
	fmt.Fprintf(w, "Upload: %d bytes in %.2f seconds = %.2f MB/s\n", bytesRead, duration, speed)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("working"))

}

func main() {
	fmt.Printf("testing dowload speed... \n")
	//TestDownloadSpeed()

	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/test", testHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

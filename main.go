package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create http client
	client := &http.Client{}

	// Create custom http request
	request, err := http.NewRequest("GET", "https://jbzd.pl", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Jbzd Imgcat - https://github.com/leszekbulawa/go_jbzd_imgcat")

	// Make http GET request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Number of bytes copied to STDOUT:", n)
}

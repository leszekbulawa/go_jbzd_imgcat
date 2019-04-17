package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
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

	// Get response body as string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	// Find title substring
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}
	// Title offset to exclude <title>
	titleStartIndex += 7

	// Find index of title closing tag
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		fmt.Println("No title closing found")
		os.Exit(0)
	}

	// Copy title do separate variable to let GC wipe whole doc
	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	// Print result
	fmt.Printf("Page title: %s\n", pageTitle)

	// Regex to find image urls
	re := regexp.MustCompile("<img src.*>")
	imageUrls := re.FindAllString(pageContent, -1)
	if imageUrls == nil {
		fmt.Println("No matches.")
	} else {
		for _, imageUrl := range imageUrls {
			fmt.Println(imageUrl)
		}
	}
}

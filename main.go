package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
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

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and print image URLs
	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			fmt.Println(imgSrc)
		}
	})


}

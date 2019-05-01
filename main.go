package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Create http client
	client := &http.Client{}

	// Get jbzd main page
	main_page := imgcat_get_request("https://jbzd.pl", client)
	defer main_page.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(main_page.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	imageUrls := []string{}

	// Find and print image URLs
	document.Find(".resource-image").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			imageUrls = append(imageUrls, imgSrc)
		}
	})

	// Seed for random int generator
	rand.Seed(time.Now().UnixNano())

	// Get image
	image := imgcat_get_request(imageUrls[rand.Intn(len(imageUrls))], client)
	defer image.Body.Close()

	// Get image content
	image_data, err := ioutil.ReadAll(image.Body)
	if err != nil{
		log.Fatal(err)
	}

	// Convert image to base64
	image_b64_data := base64.StdEncoding.EncodeToString(image_data)

	// Print image to terminal with magic header
	fmt.Printf("\033]1337;File=;size=%d;inline=1:%s%s", len(image_b64_data), image_b64_data, "\a\n")

}

func imgcat_get_request(url string, client *http.Client) *http.Response {
	// Create custom http request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Jbzd Imgcat - https://github.com/leszekbulawa/go_jbzd_imgcat")

	// Make http GET request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

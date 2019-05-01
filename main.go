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

	imageUrls := []string{}

	// Find and print image URLs
	document.Find(".resource-image").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			imageUrls = append(imageUrls, imgSrc)
		}
	})
	// Draw random image
	rand.Seed(time.Now().UnixNano())
	fmt.Println(imageUrls[rand.Intn(len(imageUrls))])

	// Get image content
	image_request, err := http.NewRequest("GET", imageUrls[rand.Intn(len(imageUrls))], nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Jbzd Imgcat - https://github.com/leszekbulawa/go_jbzd_imgcat")

	resp, err := client.Do(image_request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	image_data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}

	image_b64_data := base64.StdEncoding.EncodeToString(image_data)
	fmt.Println(len(image_b64_data))

	// Print image to terminal
	fmt.Printf("\033]1337;File=;size=%d;inline=1:%s%s", len(image_b64_data), image_b64_data, "\a\n")
}

package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://jbzd.com.pl"
const userAgent = "Jbzd Imgcat - https://github.com/leszekbulawa/go_jbzd_imgcat"

func main() {
	// Create http client
	client := &http.Client{}

	// Get jbzd main page
	mainPage := imgcatGetRequest(URL, client)
	defer mainPage.Body.Close()

	// Create a goquery document from the HTTP response
	doc, err := goquery.NewDocumentFromReader(mainPage.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	imageURLs := []string{}

	// Find and print image URLs
	doc.Find("div.article-image.article-media-image").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Find("img").Attr("src")
		if exists {
			imageURLs = append(imageURLs, src)
		}
	})

	if len(imageURLs) == 0 {
		log.Fatal("No images found")
	}

	// Choose random image
	chosenImage := imageURLs[rand.Intn(len(imageURLs))]

	// Get image
	image := imgcatGetRequest(chosenImage, client)
	defer image.Body.Close()

	// Get image content
	imageData, err := io.ReadAll(image.Body)
	if err != nil {
		log.Fatal("Error getting image content", err)
	}

	// Convert image to base64
	imageB64Data := base64.StdEncoding.EncodeToString(imageData)

	// Print image to terminal with magic header
	fmt.Printf("\033]1337;File=;size=%d;inline=1:%s%s", len(imageB64Data), imageB64Data, "\a\n")
}

func imgcatGetRequest(URL string, client *http.Client) *http.Response {
	// Create custom http request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", userAgent)

	// Make http GET request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("Request error: %s %s", URL, res.Status)
	}
	return res
}

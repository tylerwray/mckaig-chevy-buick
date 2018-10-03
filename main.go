package main

import (
	"fmt"
	"net/http"

	"github.com/tylerwray/red-scare/scraper"
)

func main() {
	baseURL := "https://www.dealerrater.com/dealer/McKaig-Chevrolet-Buick-A-Dealer-For-The-People-dealer-reviews-23685/page"
	queryParams := "filter=ONLY_POSITIVE"
	client := &http.Client{}

	s := scraper.New(baseURL, queryParams, client)

	s.Scrape(5)

	offenders := s.TopThree()

	for _, target := range offenders {
		fmt.Printf("Name: %s\n", target.Name)
		fmt.Printf("\tTitle: %s\n", target.Title)
		fmt.Printf("\tReview: %s\n\n", target.Review)
	}
}

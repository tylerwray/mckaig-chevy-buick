package scraper

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Review of McKaig Chevorlet Buick
type Review struct {
	Title  string
	Name   string
	Review string
	rank   int
}

// Scraper handles all interaction with scraping the page
type Scraper struct {
	baseURL     string
	queryParams string
	client      *http.Client
	reviews     []Review
}

// New Scraper
func New(baseURL, queryParams string, client *http.Client) *Scraper {
	return &Scraper{
		baseURL,
		queryParams,
		client,
		[]Review{},
	}
}

// Scrape stars the scraping process
func (s *Scraper) Scrape(pages int) {
	for p := 1; p <= pages; p++ {
		body, err := s.getPage(p)
		if err != nil {
			fmt.Printf("%v", err)
		}
		defer body.Close()

		err = s.parseBody(body)

		if err != nil {
			fmt.Printf("%v", err)
		}
	}
}

// TopThree all of the reviews, ranked by highest positivity
func (s *Scraper) TopThree() []Review {
	keywords := []string{
		"love",
		"perfect",
		"care",
		"honest",
		"awesome",
		"quick",
		"efficient",
		"friend",
		"best",
		"smile",
		"helpful",
		"pleasant",
		"superior",
	}

	for i, review := range s.reviews {
		haystack := strings.ToLower(review.Title + " " + review.Review)

		for _, needle := range keywords {
			if strings.Contains(haystack, needle) {
				s.reviews[i].rank = s.reviews[i].rank + 1
			}
		}
	}

	sortedReviews := sortByRank(s.reviews)

	return sortedReviews[:3]
}

type byRank []Review

func (r byRank) Len() int {
	return len(r)
}

func (r byRank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byRank) Less(i, j int) bool {
	return r[i].rank > r[j].rank
}

func sortByRank(reviews []Review) []Review {
	sort.Sort(byRank(reviews))

	return reviews
}

func (s *Scraper) getPage(page int) (io.ReadCloser, error) {
	URL := fmt.Sprintf("%s%d/?%s", s.baseURL, page, s.queryParams)
	resp, err := s.client.Get(URL)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (s *Scraper) parseBody(body io.Reader) error {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		return err
	}

	reviews := doc.Find(".review-wrapper")

	reviews.Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("h3").Text()
		name := selection.Find("span.notranslate.black").Text()
		review := selection.Find(".review-content").Text()

		s.reviews = append(s.reviews, Review{
			Title:  title[1 : len(title)-1],
			Name:   name[2:len(name)],
			Review: review,
		})
	})

	return nil
}

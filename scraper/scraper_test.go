package scraper_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tylerwray/red-scare/scraper"
)

func Test_Scraper(t *testing.T) {
	one := scraper.Review{
		Title:  "Quick and Helpful Service",
		Name:   "Ms Sunshine",
		Review: "Thankful for Jason Smith and Vernon Hall for diagnosing and fixing my Car! I always leave with a smile, they care deeply about their customers",
	}

	two := scraper.Review{
		Title:  "Like family Everytime!",
		Name:   "Pikata",
		Review: "Always love my experience with Charles and the gang.",
	}

	three := scraper.Review{
		Title:  "Helpful always!",
		Name:   "Tyler",
		Review: "Perfect help each and every time, this place is the best! I'll never go anywhere else",
	}

	var mockPageHtml = fmt.Sprintf(`
	<html>
		<div class="col-xs-12 col-sm-9 pad-none review-wrapper">
			<!-- REVIEW TITLE, USER-->
			<div class="margin-bottom-sm line-height-150">
				<h3 class="no-format inline italic-bolder font-20 dark-grey">"%s"</h3>
				<span class="italic font-18 black notranslate">- %s</span>
			</div>

			<!-- REVIEW BODY -->
			<div class="tr margin-top-md">
				<div class="td text-left valign-top ">
					<p class="font-16 review-content margin-bottom-none line-height-25">%s</p>
				</div>
			</div>
		</div>

		<div class="col-xs-12 col-sm-9 pad-none review-wrapper">
			<!-- REVIEW TITLE, USER-->
			<div class="margin-bottom-sm line-height-150">
				<h3 class="no-format inline italic-bolder font-20 dark-grey">"%s"</h3>
				<span class="italic font-18 black notranslate">- %s</span>
			</div>

			<!-- REVIEW BODY -->
			<div class="tr margin-top-md">
				<div class="td text-left valign-top ">
					<p class="font-16 review-content margin-bottom-none line-height-25">%s</p>
				</div>
			</div>
		</div>

		<div class="col-xs-12 col-sm-9 pad-none review-wrapper">
			<!-- REVIEW TITLE, USER-->
			<div class="margin-bottom-sm line-height-150">
				<h3 class="no-format inline italic-bolder font-20 dark-grey">"%s"</h3>
				<span class="italic font-18 black notranslate">- %s</span>
			</div>

			<!-- REVIEW BODY -->
			<div class="tr margin-top-md">
				<div class="td text-left valign-top ">
					<p class="font-16 review-content margin-bottom-none line-height-25">%s</p>
				</div>
			</div>
		</div>
	</html>
`, one.Title, one.Name, one.Review, two.Title, two.Name, two.Review, three.Title, three.Name, three.Review)

	queryParams := ""

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(mockPageHtml))
	}))

	defer server.Close()

	s := scraper.New(server.URL+"/", queryParams, server.Client())

	s.Scrape(1)

	list := s.TopThree()

	if len(list) != 3 {
		t.Error("Did not return the top three offenders")
	}

	if list[0].Name != one.Name {
		t.Error("Reviews are not ranked properly")
	}

	if list[1].Name != three.Name {
		t.Error("Reviews are not ranked properly")
	}

	if list[2].Name != two.Name {
		t.Error("Reviews are not ranked properly")
	}

	for _, review := range list {
		if review.Name == "" {
			t.Error("Each Review needs a name")
		}

		if review.Title == "" {
			t.Error("Each Review needs a title")
		}

		if review.Review == "" {
			t.Error("Each Review needs a review")
		}
	}
}

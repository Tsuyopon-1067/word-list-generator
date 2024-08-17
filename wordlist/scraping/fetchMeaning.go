package scraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func FetchWordMeaning(word string) string {
	c := colly.NewCollector()

	var res string
	c.OnHTML(".content-explanation", func(e *colly.HTMLElement) {
		text := e.Text
		res = strings.TrimSpace(text)
	})

	url := `https://ejje.weblio.jp/content/` + word
	err := c.Visit(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "not found"
	}
	return res
}

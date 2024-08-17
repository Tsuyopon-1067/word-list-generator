package parse

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHtml(html string) ([]Word, error) {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var res []Word

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		// ヘッダー行をスキップ
		if i == 0 {
			return
		}

		info := Word{
			Position: strings.TrimSpace(s.Find("td:nth-child(1)").Text()),
			Name:     strings.TrimSpace(s.Find("td:nth-child(2)").Text()),
			Meaning:  strings.TrimSpace(s.Find("td:nth-child(3)").Text()),
		}
		res = append(res, info)
	})

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

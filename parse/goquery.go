package parse

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

// BashOrgQuotesGQ - collects quotes (using goquery) from a bash.org quote page (i.e. http://bash.org/?browse&p=05)
func BashOrgQuotesGQ(body io.Reader) ([]string, error) {
	quotes := []string{}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return quotes, err
	}
	doc.Find("p.qt").Each(func(i int, s *goquery.Selection) {
		quotes = append(quotes, s.Text())
	})
	return quotes, nil
}

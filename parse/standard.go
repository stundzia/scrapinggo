package parse

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func findLinks(n *html.Node) []string {
	var links []string
	var linkFinder func(*html.Node)
	linkFinder = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			linkFinder(c)
		}
	}
	linkFinder(n)

	return links
}

func findQuotes(n *html.Node) []string {
	var quotes []string
	var quotesFinder func(*html.Node)
	quotesFinder = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "p" {
			for _, p := range node.Attr {
				if p.Key == "class" && p.Val == "qt" {
					quote := ""
					for n := node.FirstChild; node.NextSibling != nil; n = n.NextSibling {
						if n.Type == html.TextNode {
							quote += n.Data
						}
						if n.NextSibling == nil {
							break
						}
					}
					quotes = append(quotes, quote)
					break
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			quotesFinder(c)
		}
	}
	quotesFinder(n)

	return quotes
}

// BashOrgQuotes - collects quotes from a bash.org quote page (i.e. http://bash.org/?browse&p=05)
func BashOrgQuotes(body io.Reader) []string {
	doc, err := html.Parse(body)
	if err != nil {
		fmt.Println("err: ", err)
	}
	return findQuotes(doc)
}

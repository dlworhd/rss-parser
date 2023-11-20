package main

import (
	// "database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
)

func main() {
	wg := new(sync.WaitGroup)
	fp := gofeed.NewParser()

	rssUrls := []string{
		"https://exmaple.com/feed/",
		"https://example.com/rss.xml",
	}

	wg.Add(len(rssUrls))

	for _, rss := range rssUrls {
		go proceedFeed(fp, rss, wg)
	}

	wg.Wait()

}

func proceedFeed(fp *gofeed.Parser, rss string, wg *sync.WaitGroup) {
	innerWg := new(sync.WaitGroup)

	defer wg.Done()

	feeds, err := fp.ParseURL(rss)
	if err != nil {
		log.Printf("Error parsing URL %s: %v", rss, err)
		return
	}

	innerWg.Add(len(feeds.Items))

	for _, feed := range feeds.Items {
		go feedItem(feed, innerWg)
	}

	innerWg.Wait()
}

func fetchHtml(url string) (string, error) {
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(content), nil

}

func extractThumbnail(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	var thumbnailURL string
	var traverseFunc func(*html.Node)
	traverseFunc = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			for _, attr := range n.Attr {
				if attr.Key == "property" && attr.Val == "og:image" {
					for _, subAttr := range n.Attr {
						if subAttr.Key == "content" {
							thumbnailURL = subAttr.Val
							return
						}
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseFunc(c)
		}
	}
	traverseFunc(doc)

	return thumbnailURL, nil
}

func feedItem(feed *gofeed.Item, innerWg *sync.WaitGroup) {
	defer innerWg.Done()
	fg := feed.GUID
	ft := feed.Title
	fd := feed.Description
	fl := feed.Link
	fp := feed.Published

	html, err := fetchHtml(fl)
	if err != nil {
		log.Printf("Error fetching HTML for %s: %v", fl, err)
		return
	}

	th, err := extractThumbnail(html)
	if err != nil {
		log.Printf("Error extracting thumbnail for %s: %v", fl, err)
		return
	}

	fmt.Println("GUID: ", fg)
	fmt.Println("Title: ", ft)
	fmt.Println("Description: ", fd)
	fmt.Println("Link: ", fl)
	fmt.Println("Thumbnail: ", th)
	fmt.Println("Date: ", fp)
	
}

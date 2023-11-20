package rss_parser

import (
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
)

type Feed struct {
	guid        string
	title       string
	thumbnail   string
	description string
	link        string
	published   string
}

// func main() {
// 	wg := new(sync.WaitGroup)
// 	fp := gofeed.NewParser()

// 	rssUrls := []string{
// 		"https://toss.tech/rss.xml",
// 	}

// 	wg.Add(len(rssUrls))

// 	var outerFeedArr [][]Feed

// 	for range rssUrls {
// 		outerFeedArr = append(outerFeedArr, nil)
// 	}

// 	for i, rss := range rssUrls {
// 		ProceedFeed(fp, rss, wg, outerFeedArr, i)
// 	}

// 	for _, innerFeedArr := range outerFeedArr {
// 		for _, feed := range innerFeedArr {
// 			fmt.Println("GUID: ", feed.guid)
// 			fmt.Println("Title: ", feed.title)
// 			fmt.Println("Thumbnail: ", feed.thumbnail)
// 			fmt.Println("Description : ", feed.description)
// 			fmt.Println("Link: ", feed.link)
// 			fmt.Println("Published: ", feed.published)
// 			fmt.Printf("\n")
// 		}
// 	}

// 	wg.Wait()

// }

func ProceedFeed(fp *gofeed.Parser, rss string, wg *sync.WaitGroup, outerFeedArr [][]Feed, outIndex int) {
	innerWg := new(sync.WaitGroup)

	defer wg.Done()

	feeds, err := fp.ParseURL(rss)
	if err != nil {
		log.Printf("Error parsing URL %s: %v", rss, err)
		panic(err)
	}

	innerWg.Add(len(feeds.Items))

	var innerFeedArr []Feed

	for _, feed := range feeds.Items {
		go AddFeedItem(feed, innerWg, &innerFeedArr)
	}

	innerWg.Wait()
	(outerFeedArr)[outIndex] = innerFeedArr

}

func FetchHtml(url string) (string, error) {
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

func ExtractThumbnail(htmlContent string) (string, error) {
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

func AddFeedItem(feed *gofeed.Item, innerWg *sync.WaitGroup, feedArr *[]Feed) {
	defer innerWg.Done()
	html, err := FetchHtml(feed.Link)
	if err != nil {
		log.Printf("Error fetching HTML for %s: %v", feed.Link, err)
		panic(err)
	}
	if err != nil {
		log.Printf("Error extracting thumbnail for %s: %v", feed.Link, err)
		panic(err)
	}
	thumbnail, _ := ExtractThumbnail(html)

	f := Feed{
		guid:        feed.GUID,
		title:       feed.Title,
		thumbnail:   thumbnail,
		description: feed.Description,
		link:        feed.Link,
		published:   feed.Published,
	}

	*feedArr = append(*feedArr, f)
}

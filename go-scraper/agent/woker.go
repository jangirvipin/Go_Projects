package agent

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func Worker(urlChan <-chan string, resultChan chan<- *Article, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range urlChan {
		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			select {
			case <-time.After(30 * time.Second):
				fmt.Println("Timeout")
				cancel()
			case <-ctx.Done():
				fmt.Println("Scrapping completed")
			}
		}()
		article := scrapeArticle(url, ctx)
		resultChan <- article
		cancel()
	}

}

func scrapeArticle(url string, ctx context.Context) *Article {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()
	fmt.Println("Title:", title)

	var content strings.Builder

	doc.Find("div.article-body-commercial-selector p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		content.WriteString(text)
	})

	return &Article{
		URL:     url,
		Title:   title,
		Content: content.String(),
	}
}

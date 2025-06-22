package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jangirvipin/go-scraper/agent"
	"github.com/jangirvipin/go-scraper/parse"
	"log"
	"net/http"
)

var links []string

func main() {

	res, err := http.Get("https://www.theguardian.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			if parse.ValidLinksOnly(link) {
				final := parse.Normalize(link)
				links = append(links, final)
			}
		}
	})

	ag := agent.Agent{
		WorkerCount: 5,
		URLS:        links,
	}
	
	fmt.Print("Agent started")
	ag.Agent()

}

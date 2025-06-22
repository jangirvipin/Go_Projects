package agent

import (
	"fmt"
	"sync"
)

type Article struct {
	URL     string
	Title   string
	Content string
}
type Agent struct {
	WorkerCount int
	URLS        []string
	Articles    []*Article
}

func (a *Agent) Agent() {
	var wg sync.WaitGroup
	urlChan := make(chan string)
	resultChan := make(chan *Article)

	for i := 0; i < a.WorkerCount; i++ {
		wg.Add(1)
		fmt.Println("Worker started")
		go Worker(urlChan, resultChan, &wg)
	}

	go func() {
		for _, url := range a.URLS {
			urlChan <- url
		}
		close(urlChan)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		a.Articles = append(a.Articles, result)
		fmt.Println("Scrapped")
	}
}

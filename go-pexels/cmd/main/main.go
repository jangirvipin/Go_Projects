package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}
type Photo struct {
	ID              string `json:"id"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	URL             string `json:"url"`
	Photographer    string `json:"photographer"`
	PhotographerURL string `json:"photographer_url"`
	PhotographerID  int    `json:"photographer_id"`
	AvgColor        string `json:"avg_color"`
	SRC             source `json:"src"`
	Liked           bool   `json:"liked"`
	ALT             string `json:"alt"`
}

type source struct {
	Original  string `json:"original"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
	Medium    string `json:"medium"`
	Large2X   string `json:"large2x"`
	Large     string `json:"large"`
}

type SearchResult struct {
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	TotalResults int     `json:"total_results"`
	NextPage     string  `json:"next_page"`
	PrevPage     string  `json:"prev_page"`
	Photos       []Photo `json:"photos"`
}

func (c *Client) GetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	res, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", res.Status)
	}
	time, err := strconv.Atoi(res.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return nil, fmt.Errorf("error parsing rate limit: %v", err)
	} else {
		c.RemainingTimes = int32(time)
	}
	return res, nil
}

func (c *Client) SearchPhotos(query string, page int, perPage int) (*SearchResult, error) {
	url := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s&page=%d&per_page=%d", query, page, perPage)
	res, err := c.GetRequest(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	err = json.Unmarshal(data, &result)
	return &result, err
}

func (c *Client) CuratedPhotos(page int, perPage int) (*SearchResult, error) {
	url := fmt.Sprintf("https://api.pexels.com/v1/curated?page=%d&per_page=%d", page, perPage)
	res, err := c.GetRequest(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetPhoto(id string) (*Photo, error) {
	url := fmt.Sprintf("https://api.pexels.com/v1/photos/%s", id)
	res, err := c.GetRequest(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var photo Photo
	err = json.Unmarshal(data, &photo)
	return &photo, err
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c, RemainingTimes: 0}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	client := NewClient(os.Getenv("TOKEN"))
	result, err := client.SearchPhotos("nature", 1, 10)
	fmt.Println(result, err)
}

package service

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"golang-hatena-client/model"
)

type HatenaRssReader interface {
	Read(url string)
}

type HatenaClient struct {
	HatenaRssReader
}

func (h *HatenaClient) Read(url string) (*model.HatenaFeed, error) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error: rss feed get error. err: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	var feed model.HatenaFeed
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		fmt.Printf("Error: xml unmarshal error. err:  %v", err)
		return nil, err
	}
	return &feed, err
}

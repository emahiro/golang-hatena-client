package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

	body := resp.Body
	defer body.Close()

	b, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Printf("Error: read error. err: %v", err)
		return nil, err
	}

	var feed model.HatenaFeed
	if err := xml.Unmarshal(b, &feed); err != nil {
		fmt.Printf("Error: xml unmarshal error. err:  %v", err)
		return nil, err
	}

	return &feed, err
}

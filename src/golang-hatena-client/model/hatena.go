package model

import (
	"encoding/xml"
)

type HatenaFeed struct {
	XMLName         xml.Name         `xml:"RDF"`
	Title           string           `xml:"channel>title"`
	Link            string           `xml:"channel>link"`
	Description     string           `xml:"channel>description"`
	HatenaBookmarks []HatenaBookmark `xml:"item"`
}

type HatenaBookmark struct {
	XMLName       xml.Name `xml:"item"`
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Content       string   `xml:"content"`
	Date          string   `xml:"date"`
	Subject       string   `xml:"subject"`
	BookmarkCount int64    `xml:"bookmarkcount"`
}

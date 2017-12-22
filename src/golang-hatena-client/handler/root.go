package handler

import (
	"fmt"
	"net/http"

	"golang-hatena-client/service"

	"github.com/k0kubun/pp"
)

func Top(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func GetRSS(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	url := v.Get("url")
	if url == "" {
		fmt.Printf("Error: %v", pp.Sprintf("%v", "rss url not setted"))
	}

	// hatenaが含まれない場合も処理進ませない

	client := service.HatenaClient{}
	feed, err := client.Read(url)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	fmt.Printf("%v", pp.Sprint(feed))
}

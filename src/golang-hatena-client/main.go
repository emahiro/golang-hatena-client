package main

import (
	"fmt"
	"net/http"

	"golang-hatena-client/web"
)

var port = 8080

func main() {
	router := web.Router()
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}

package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	url_example := "https%3A%2F%2Fwww.google.com%2F"
	a, err := url.QueryUnescape(url_example)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(a)
}

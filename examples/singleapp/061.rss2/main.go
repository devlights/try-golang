package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type (
	Item struct {
		XMLName xml.Name `xml:"item"`
		Title   string   `xml:"title"`
	}
	Rss struct {
		XMLName xml.Name `xml:"rss"`
		Items   []Item   `xml:"channel>item"`
	}
)

const (
	URL = "https://devlights.hatenablog.com/rss"
)

func panicOnErr[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// ブログ [いろいろ備忘録日記](https://devlights.hatenablog.com/) から RSS 2.0 を取得して表示するサンプルです。
//
// REFERENCES:
//   - https://qiita.com/you8/items/e903fd463cf770688e1e
func main() {
	fmt.Printf("[URL] %s\n", URL)

	resp := panicOnErr(http.Get(URL))
	body := panicOnErr(io.ReadAll(resp.Body))
	defer resp.Body.Close()

	var rss Rss
	if err := xml.Unmarshal(body, &rss); err != nil {
		panic(err)
	}

	for i, item := range rss.Items {
		fmt.Printf("[%2d] %s\n", i+1, item.Title)
	}
}

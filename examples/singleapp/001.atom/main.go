package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type (
	Feed struct {
		XMLName xml.Name `xml:"feed"`
		Titles  []string `xml:"entry>title"`
	}
)

const (
	URL = "https://devlights.hatenablog.com/feed"
)

func panicOnErr[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// ブログ [いろいろ備忘録日記](https://devlights.hatenablog.com/) から Atom Feed を取得して表示するサンプルです。
//
// REFERENCES:
//   - https://qiita.com/you8/items/e903fd463cf770688e1e
func main() {
	fmt.Printf("[URL] %s\n", URL)

	resp := panicOnErr(http.Get(URL))
	body := panicOnErr(io.ReadAll(resp.Body))
	defer resp.Body.Close()

	var feed Feed
	if err := xml.Unmarshal(body, &feed); err != nil {
		panic(err)
	}

	for i, title := range feed.Titles {
		fmt.Printf("[%2d] %s\n", i+1, title)
	}
}

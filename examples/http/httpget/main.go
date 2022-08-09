// http.Get() を使って HTTP GET リクエストを試してみるサンプルです。
//
// リクエストの発行先は JSONPlaceholder (https://jsonplaceholder.typicode.com/) を
// 使わせてもらっています。
//
// REFERENCES:
//   - https://dev.to/devkiran/make-an-http-get-request-in-go-58gf
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (me *Post) String() string {
	return fmt.Sprintf("Uid: %d, Id: %d, Title: %s..., Body: %s...", me.UserId, me.Id, me.Title[:5], me.Body[:5])
}

const (
	url = "https://jsonplaceholder.typicode.com/posts/1"
)

var (
	appLog = log.New(os.Stderr, "", 0)
	errLog = log.New(os.Stderr, "[Error] ", 0)
)

func main() {
	// -------------------------------------------
	// Send http GET request
	// -------------------------------------------

	res, err := http.Get(url)
	if err != nil {
		errLog.Println(err)
		return
	}
	defer res.Body.Close()

	// -------------------------------------------
	// Check http status code
	// -------------------------------------------

	if res.StatusCode != http.StatusOK {
		errLog.Printf("http status code: %d", res.StatusCode)
		return
	}

	// -------------------------------------------
	// Decode response to JSON
	// -------------------------------------------

	var (
		post    = &Post{}
		decoder = json.NewDecoder(res.Body)
	)

	err = decoder.Decode(post)
	if err != nil {
		errLog.Println(err)
		return
	}

	// -------------------------------------------
	// Show results
	// -------------------------------------------

	appLog.Printf("status: %d, resonse: %s\n", res.StatusCode, post)
}

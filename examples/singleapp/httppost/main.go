//
// http.NewRequest() を使って HTTP POST リクエストを試してみるサンプルです。
//
// リクエストの発行先は JSONPlaceholder (https://jsonplaceholder.typicode.com/) を
// 使わせてもらっています。
//
// REFERENCES:
//   - https://dev.to/devkiran/make-an-http-post-request-in-go-29p
//
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type (
	Post struct {
		UserId int    `json:"userId"`
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}
)

func (me *Post) String() string {
	return fmt.Sprintf("Uid: %d, Id: %d, Title: %s, Body: %s", me.UserId, me.Id, me.Title, me.Body)
}

const (
	method      = "POST"
	url         = "https://jsonplaceholder.typicode.com/posts"
	contentType = "application/json"
)

var (
	appLog = log.New(os.Stdout, "", 0)
	errLog = log.New(os.Stderr, "[Error] ", 0)
)

func main() {
	// -------------------------------------------
	// Create request and set headers
	// -------------------------------------------

	var (
		body = []byte(`{"userId": 999, "title": "try-golang", "body": "try-golang/cmd/httppost"}`)
		buf  = bytes.NewBuffer(body)
	)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		errLog.Println(err)
		return
	}
	req.Header.Add("Content-Type", contentType)

	// -------------------------------------------
	// Send http POST request
	// -------------------------------------------

	var (
		client = &http.Client{}
	)

	res, err := client.Do(req)
	if err != nil {
		errLog.Println(err)
		return
	}
	defer res.Body.Close()

	// -------------------------------------------
	// Check http status code
	// -------------------------------------------

	if res.StatusCode != http.StatusCreated {
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

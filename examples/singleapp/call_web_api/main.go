// 無償で利用できる Web API である Free Weather API を使って、HTTPリクエストを送信し
// 結果データとして受信した JSON を表示するサンプルです。
//
// # REFERENCES
//   - https://open-meteo.com/
//   - https://open-meteo.com/en/docs
//   - https://paiza.hatenablog.com/entry/2021/11/04/130000
//   - https://tech.yappli.io/entry/go_unmarshal_interface
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type (
	JsonData struct {
		Daily Daily `json:"daily"`
	}

	Daily struct {
		Time []string  `json:"time"`
		Max  []float32 `json:"temperature_2m_max"`
		Min  []float32 `json:"temperature_2m_min"`
	}
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	endpoint := url.URL{
		Scheme:   "https",
		Host:     "api.open-meteo.com",
		Path:     "v1/forecast",
		RawQuery: "latitude=34.69&longitude=135.50&daily=temperature_2m_max,temperature_2m_min&timezone=Asia%2FTokyo", // OSAKA
	}

	// Get
	resp, err := http.Get(endpoint.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Decode
	data := JsonData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	// Show
	daily := data.Daily
	for i := 0; i < len(daily.Time); i++ {
		fmt.Printf("Time: %s\tMax: %.1f\tMin: %.1f\n", daily.Time[i], daily.Max[i], daily.Min[i])
	}

	return nil
}

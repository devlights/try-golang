package base64op

import (
	"bytes"
	"io"
	"net/http"
)

type (
	downloder string
)

func (me downloder) download() (*bytes.Buffer, error) {
	var (
		client http.Client
		resp   *http.Response
		buf    = new(bytes.Buffer)
		err    error
	)

	if resp, err = client.Get(string(me)); err != nil {
		return buf, err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(buf, resp.Body); err != nil {
		return buf, err
	}

	return buf, nil
}

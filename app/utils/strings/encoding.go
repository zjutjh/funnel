package strings

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
	"io"
)

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, _ := charset.DetermineEncoding(data, ""); len(name) != 0 {
			return name
		}
	}

	return "utf-8"
}

func DecodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
	if charset == "" {
		charset = detectContentCharset(body)
	}

	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}

	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}

	return body, nil
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var recommendRe = regexp.MustCompile(`<h2>[\s\S]*?</h2>`)
var htmlTagRe = regexp.MustCompile("<[\\s\\S]*?>")

func main() {
	url := "https://www.thepaper.cn/"

	body, err := Fetch(url)
	if err != nil {
		fmt.Println("read content failed:", err)

		return
	}

	matches := recommendRe.FindAllSubmatch(body, -1)
	for _, m := range matches {
		c := htmlTagRe.ReplaceAllString(string(m[0]), "")

		fmt.Println("fetch card news:", c)
	}
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %v\n", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return io.ReadAll(utf8Reader)
}

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	peek, err := r.Peek(1024)
	if err != nil {
		fmt.Println("fetch err:", err)

		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(peek, "")

	return e
}

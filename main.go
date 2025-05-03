package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	url := "https://www.thepaper.cn/"

	body, err := Fetch(url)
	if err != nil {
		fmt.Println("read content failed:", err)

		return
	}

	doc, err := htmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Println("htmlquery.Parse failed:", err)
	}

	nodes := htmlquery.Find(doc, `//div[@class='small_toplink__GmZhY']/a[@target='_blank']/h2`)
	for _, n := range nodes {

		fmt.Println("fetch card news:", n.FirstChild.Data)
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

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("fetch url error:", err)

		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("Error status code: %v\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response body error:", err)

		return
	}

	numsLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has: %d links!\n", numsLinks)

	exist := strings.Contains(string(body), "黄金")
	fmt.Printf("是否存在黄金:%v\n", exist)

	exist = bytes.Contains(body, []byte("黄金"))
	fmt.Printf("是否存在黄金:%v\n", exist)
}

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/miaoerwu/crawler/collect"
)

func main() {
	url := "https://book.douban.com/subject/1007305/"

	fetch := collect.BrowserFetch{
		Timeout: 2000 * time.Millisecond,
	}
	body, err := fetch.Get(url)
	if err != nil {
		fmt.Println("read body failed:", err)

		return
	}

	fmt.Println(string(body))

	fmt.Println()
	fmt.Println("------------------------------------------")
	fmt.Println()

	ctx, cancelFunc := chromedp.NewContext(context.Background())
	defer cancelFunc()

	ctx, cancelFunc = context.WithTimeout(ctx, 15*time.Second)
	defer cancelFunc()

	var example string
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://pkg.go.dev/time"),
		chromedp.WaitVisible("body > footer"),
		chromedp.Click("#example-After", chromedp.NodeVisible),
		chromedp.Value("#example-After textarea", &example),
	)
	if err != nil {
		fmt.Println("read body failed:", err)

		return
	}

	fmt.Println(example)
}

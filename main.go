package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"

	"github.com/miaoerwu/crawler/collect"
	"github.com/miaoerwu/crawler/proxy"
)

func main() {
	urls := []string{"http://127.0.0.1:8888", "http://127.0.0.1:8889"}
	p, err := proxy.NewRoundRobinSwitcher(urls...)
	if err != nil {
		fmt.Println("RoundRobinProxySwitcher failed")

		return
	}

	url := "https://google.com"
	fetch := collect.BrowserFetch{
		Timeout: 2000 * time.Millisecond,
		Proxy:   p,
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

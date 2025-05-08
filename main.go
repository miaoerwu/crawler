package main

import (
	"context"
	"io"
	"time"

	"github.com/chromedp/chromedp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/miaoerwu/crawler/collect"
	"github.com/miaoerwu/crawler/log"
)

func main() {
	plugin, closer := log.NewFilePlugin("./log.txt", zapcore.InfoLevel)
	defer func(closer io.Closer) {
		_ = closer.Close()
	}(closer)
	logger := log.NewLogger(plugin)
	logger.Info("log init end")

	//urls := []string{"http://127.0.0.1:8888", "http://127.0.0.1:8889"}
	//p, err := proxy.NewRoundRobinSwitcher(urls...)
	//if err != nil {
	//	logger.Error("RoundRobinProxySwitcher failed")
	//
	//	return
	//}

	url := "https://google.com"
	fetch := collect.BrowserFetch{
		Timeout: 2000 * time.Millisecond,
		//Proxy:   p,
	}
	body, err := fetch.Get(url)
	if err != nil {
		logger.Error("read body failed:", zap.Error(err))

		return
	}

	logger.Info(string(body))

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
		logger.Error("read body failed:", zap.Error(err))

		return
	}

	logger.Info(example)
}

package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/miaoerwu/crawler/collect"
	"github.com/miaoerwu/crawler/log"
	"github.com/miaoerwu/crawler/parse/doubangroup"
)

func main() {
	plugin := log.NewStdoutPlugin(zapcore.InfoLevel)
	logger := log.NewLogger(plugin)
	logger.Info("log init end")

	//urls := []string{"http://127.0.0.1:8888", "http://127.0.0.1:8889"}
	//p, err := proxy.NewRoundRobinSwitcher(urls...)
	//if err != nil {
	//	logger.Error("RoundRobinProxySwitcher failed")
	//
	//	return
	//}

	cookie := ""
	var worklist []*collect.Request
	for i := 0; i <= 0; i += 25 {
		url := fmt.Sprintf("https://www.douban.com/group/szsh/discussion?start=%d", i)

		worklist = append(worklist, &collect.Request{
			Url:       url,
			Cookies:   cookie,
			ParseFunc: doubangroup.ParseUrl,
		})
	}

	fetch := collect.BrowserFetch{
		Timeout: 2000 * time.Millisecond,
		//Proxy:   p,
	}

	for len(worklist) > 0 {
		size := len(worklist)

		for range size {
			req := worklist[0]
			worklist = worklist[1:]

			body, err := fetch.Get(req)
			if err != nil {
				logger.Error("read body failed:", zap.Error(err))

				continue
			}

			res := req.ParseFunc(body, req)
			for _, item := range res.Items {
				logger.Info("result", zap.String("get url:", item.(string)))
			}
			worklist = append(worklist, res.Requests...)
		}
	}

	//logger.Info(string(body))
	//
	//ctx, cancelFunc := chromedp.NewContext(context.Background())
	//defer cancelFunc()
	//
	//ctx, cancelFunc = context.WithTimeout(ctx, 15*time.Second)
	//defer cancelFunc()
	//
	//var example string
	//err = chromedp.Run(ctx,
	//	chromedp.Navigate("https://pkg.go.dev/time"),
	//	chromedp.WaitVisible("body > footer"),
	//	chromedp.Click("#example-After", chromedp.NodeVisible),
	//	chromedp.Value("#example-After textarea", &example),
	//)
	//if err != nil {
	//	logger.Error("read body failed:", zap.Error(err))
	//
	//	return
	//}
	//
	//logger.Info(example)
}

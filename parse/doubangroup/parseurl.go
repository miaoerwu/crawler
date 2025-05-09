package doubangroup

import (
	"regexp"

	"github.com/miaoerwu/crawler/collect"
)

const urlListRe = `(?si)` + // 启用多行和忽略大小写
	`<a\s+` + // 匹配 <a 标签开始
	`href="(https://www\.douban\.com/group/topic/\d+/[^"]*)"` + // 完整 href（包含数字和参数，但不拆分）
	`\s+title="([^"]*)"` + // 固定顺序的 title
	`[^>]*>` + // 跳过其他属性
	`\s*(.*?)\s*</a>` // 捕获内容并自动去空格

func ParseUrl(content []byte, req *collect.Request) collect.ParseResult {
	re := regexp.MustCompile(urlListRe)

	matches := re.FindAllSubmatch(content, -1)

	result := collect.ParseResult{}

	for _, m := range matches {
		u := string(m[1])
		result.Requests = append(result.Requests, &collect.Request{
			Url:     u,
			Cookies: req.Cookies,
			ParseFunc: func(c []byte, request *collect.Request) collect.ParseResult {
				return GetContent(c, u)
			},
		})
	}

	return result
}

const contentRe = `<div class="topic-content">[\s\S]*?阳台[\s\S]*?<div`

func GetContent(contents []byte, url string) collect.ParseResult {
	re := regexp.MustCompile(contentRe)

	if !re.Match(contents) {
		return collect.ParseResult{
			Items: []any{},
		}
	}

	return collect.ParseResult{
		Items: []any{url},
	}
}

package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// 实例化默认采集器
	c := colly.NewCollector()
	// 在发出请求之前将带有“url”键的URL放入请求的上下文中
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})
	// 在发出请求后从请求的上下文中获取“url”
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Ctx.Get("url"))
	})
	// 从https://en.wikipedia.org开始爬取
	c.Visit("https://en.wikipedia.org/")
}

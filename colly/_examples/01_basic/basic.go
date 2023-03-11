package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// 实例化默认收集器
	c := colly.NewCollector(
		// 访问的域仅限于：hackerspaces.org 和 wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)
	// 对于每个具有 href 属性的 a 元素调用回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// 打印链接
		fmt.Printf("Link found: %!q(MISSING) -> %!s(MISSING)\n", e.Text, link)
		// 访问在页面中找到的链接
		// 只有在 AllowedDomains 中的链接才会被访问
		c.Visit(e.Request.AbsoluteURL(link))
	})
	// 在发出请求之前打印 "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// 从 https://hackerspaces.org 开始爬取
	c.Visit("https://hackerspaces.org/")
}

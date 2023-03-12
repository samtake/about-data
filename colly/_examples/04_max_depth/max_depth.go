package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// 创建一个最大深度为1的Collector对象
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	// 在每个具有href属性的a元素上调用回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// 输出链接
		fmt.Println(link)
		// 访问链接所指向的页面
		e.Request.Visit(link)
	})
	// 开始爬取维基百科的数据
	c.Visit("https://en.wikipedia.org/")
}

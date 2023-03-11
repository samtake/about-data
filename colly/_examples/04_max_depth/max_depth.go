package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(1),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Println(link)
		// Visit link found on page
		e.Request.Visit(link)
	})

	// Start scraping on https://en.wikipedia.org
	c.Visit("https://en.wikipedia.org/")
}

/**
MaxDepth 是 colly 框架中的一个选项，表示最大的访问深度。
当设置 MaxDepth 为 1 时，只有当前页面中的链接会被访问，不会继续跟进这些链接所指向的页面中的链接，
也就是说只会在当前页面进行抓取，不会进一步抓取其他页面。
因此，上述语句的意思是：由于 MaxDepth 选项被设置为 1，所以只有当前爬取页面的链接会被访问，不会继续访问其他页面的链接。
**/

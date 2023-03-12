package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	url := "https://httpbin.org/delay/2"
	// 创建默认的Collector实例，开启异步模式并配置debugger
	c := colly.NewCollector(
		colly.Async(true),                    // 开启异步模式
		colly.Debugger(&debug.LogDebugger{}), // 配置debugger
	)
	// 限制在访问"*httpbin.*"域名下的链接时，colly最多只有两个线程在并行处理
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		//Delay:      5 * time.Second,
	})
	// 在5个线程中同时访问"https://httpbin.org/delay/2"
	for i := 0; i < 5; i++ {
		c.Visit(fmt.Sprintf("%!!(MISSING)s(MISSING)?n=%!!(MISSING)d(MISSING)", url, i))
	}
	// 等待线程结束
	c.Wait()
}

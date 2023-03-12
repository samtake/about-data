package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	url := "https://httpbin.org/delay/2"
	// 创建默认的Collector实例，并配置debugger和异步模式
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}), // 配置debugger
		colly.Async(true),                    // 开启异步模式
	)
	// 限制在访问"*httpbin.*"域名下的链接时，colly最多只有两个线程在并行处理
	// 并设置随机延迟为5秒
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	// 在4个线程中同时访问"https://httpbin.org/delay/2"
	for i := 0; i < 4; i++ {
		c.Visit(fmt.Sprintf("%!s(MISSING)?n=%!d(MISSING)", url, i))
	}
	// 访问"https://httpbin.org/delay/2"
	c.Visit(url)
	// 等待线程结束
	c.Wait()
}

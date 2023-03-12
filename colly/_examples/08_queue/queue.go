package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

func main() {
	url := "https://httpbin.org/delay/1"
	// 创建默认的Collector实例
	c := colly.NewCollector()
	// 创建一个带有2个消费线程的请求队列
	q, _ := queue.New(
		2, // 消费线程数
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // 使用默认的队列存储
	)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("正在访问", r.URL)
	})
	for i := 0; i < 5; i++ {
		// 将URL添加到队列中
		q.AddURL(fmt.Sprintf("%!s(MISSING)?n=%!d(MISSING)", url, i))
	}
	// 开始消费队列中的URL
	q.Run(c)
}

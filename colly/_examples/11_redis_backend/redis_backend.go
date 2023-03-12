package main

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/gocolly/redisstorage"
)

func main() {
	urls := []string{
		"http://httpbin.org/",
		"http://httpbin.org/ip",
		"http://httpbin.org/cookies/set?a=b&c=d",
		"http://httpbin.org/cookies",
	}
	c := colly.NewCollector()
	// 创建Redis存储
	storage := &redisstorage.Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "httpbin_test",
	}
	// 将存储添加到Collector实例中
	err := c.SetStorage(storage)
	if err != nil {
		panic(err)
	}
	// 清空之前的存储数据
	if err := storage.Clear(); err != nil {
		log.Fatal(err)
	}
	// 关闭Redis客户端
	defer storage.Client.Close()
	// 使用Redis存储创建一个新的请求队列
	q, _ := queue.New(2, storage)
	c.OnResponse(func(r *colly.Response) {
		log.Println("Cookies:", c.Cookies(r.Request.URL.String()))
	})
	// 将URL添加到队列中
	for _, u := range urls {
		q.AddURL(u)
	}
	// 开始请求队列中的请求
	q.Run(c)
}

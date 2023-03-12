package main

import (
	"github.com/gocolly/colly"
	"log"
)

func main() {
	// 创建一个新的Collector对象
	c := colly.NewCollector()
	// 认证登录
	err := c.Post("http://example.com/login", map[string]string{"username": "admin", "password": "admin"})
	if err != nil {
		log.Fatal(err)
	}
	// 登录后，在响应中添加回调函数
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})
	// 开始爬取目标网站的数据
	c.Visit("https://example.com/")
}

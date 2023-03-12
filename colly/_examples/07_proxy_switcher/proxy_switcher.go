package main

import (
	"bytes"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func main() {
	// 创建了一个默认的Collector实例，并允许URL重访问
	c := colly.NewCollector(colly.AllowURLRevisit())

	// 创建了一个包含两个socks5代理的代理轮换器，并将其设置为Collector实例的代理函数
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	// Fetch httpbin.org/ip five times
	for i := 0; i < 5; i++ {
		c.Visit("https://httpbin.org/ip")
	}
}

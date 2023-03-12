package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// 定义一个 colly extension，用于处理数据并返回结果
type MyExtension struct {
	Data []string
}

// 实现 Extension 接口的方法，用于处理数据
func (m *MyExtension) OnResponse(response *colly.Response) error {
	// 将爬取到的数据存储到 MyExtension.Data 中
	m.Data = append(m.Data, string(response.Body))
	return nil
}
func main() {
	c := colly.NewCollector()
	ext := &MyExtension{}
	// 添加 colly extension
	c.SetExtensions(ext)
	// 使用 colly extension 进行数据处理
	c.Visit("https://www.example.com/")
	// 输出处理结果
	fmt.Printf("Data: %!v(MISSING)\n", ext.Data)
}

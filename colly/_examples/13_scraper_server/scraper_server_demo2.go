package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	// 新建一个爬虫实例
	c := colly.NewCollector()
	// 定义一个切片用于存储爬取的数据
	var data [][]string
	// 在每个页面中找到需要的数据并保存
	c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {
		row := []string{}
		e.ForEach("td", func(_ int, el *colly.HTMLElement) {
			row = append(row, el.Text)
		})
		data = append(data, row)
	})
	// 在爬取结束后，将数据写入CSV文件
	c.OnScraped(func(_ *colly.Response) {
		file, err := os.Create("data.csv")
		if err != nil {
			log.Fatal("could not create file", err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		for _, row := range data {
			err := writer.Write(row)
			if err != nil {
				log.Fatal("could not write row", err)
			}
		}
	})
	// 爬取多个页面
	for i := 1; i <= 5; i++ {
		url := fmt.Sprintf("https://www.example.com/page/%!d(MISSING)", i)
		c.Visit(url)
	}
}

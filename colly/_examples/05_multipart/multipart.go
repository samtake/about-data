package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gocolly/colly"
)

// 生成包含表单数据的map
func generateFormData() map[string][]byte {
	f, _ := os.Open("gocolly.jpg")
	defer f.Close()
	imgData, _ := ioutil.ReadAll(f)
	return map[string][]byte{
		"firstname": []byte("one"),
		"lastname":  []byte("two"),
		"email":     []byte("onetwo@example.com"),
		"file":      imgData,
	}
}

// 创建HTTP服务器并等待POST请求
func setupServer() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request")
		err := r.ParseMultipartForm(10000000)
		if err != nil {
			fmt.Println("server: Error")
			w.WriteHeader(500)
			w.Write([]byte("<html><body>Internal Server Error</body></html>"))
			return
		}
		w.WriteHeader(200)
		fmt.Println("server: OK")
		w.Write([]byte("<html><body>Success</body></html>"))
	}
	go http.ListenAndServe(":8080", handler)
}
func main() {
	// 创建一个包含表单数据的map
	formData := generateFormData()
	// 创建HTTP服务器并等待POST请求
	setupServer()
	// 创建一个Colly Collector对象
	c := colly.NewCollector(colly.AllowURLRevisit(), colly.MaxDepth(5))
	// 在每个a元素上调用回调函数
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// 输出每个HTML页面的文本
		fmt.Println(e.Text)
		// 延迟1秒
		time.Sleep(1 * time.Second)
		// 将表单数据作为multipart/form-data格式的POST请求发送到HTTP服务器
		e.Request.PostMultipart("http://localhost:8080/", formData)
	})
	// 在发送请求之前输出"Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Posting gocolly.jpg to", r.URL.String())
	})
	// 开始爬取
	c.PostMultipart("http://localhost:8080/", formData)
	c.Wait()
}

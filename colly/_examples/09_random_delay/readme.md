
随机延时其实是更好的模拟人为的用户操作
```bash
// 限制在访问"*httpbin.*"域名下的链接时，colly最多只有两个线程在并行处理
// 并设置随机延迟为5秒
c.Limit(&colly.LimitRule{
	DomainGlob:  "*httpbin.*",
	Parallelism: 2,
	RandomDelay: 5 * time.Second,
})
```
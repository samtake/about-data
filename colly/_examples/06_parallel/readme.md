

`c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})`

- `DomainGlob: "*"`设置并发限制的域名范围为所有域名
- `Parallelism: 2`最大并发数为2
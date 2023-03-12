在这个示例中，我们定义了一个名为 MyExtension 的 colly extension，用于处理爬取到的数据。MyExtension 实现了 Extension 接口的 OnResponse 方法，用于处理响应数据。在这个示例中，我们将爬取到的数据存储到 MyExtension.Data 中。然后，我们创建一个 colly 爬虫实例 c，并使用 c.SetExtensions 方法将 MyExtension 添加到爬虫实例中。最后，我们使用 c.Visit 方法爬取数据，并输出处理结果。
需要注意的是，在实际使用 colly extension 时，需要根据具体的需求来实现 Extension 接口的各个方法。同时，还需要考虑如何将爬取到的数据传递给 colly extension。一般来说，可以使用 colly 爬虫实例的 Context 属性来传递数据。

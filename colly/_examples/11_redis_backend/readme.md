
- 首先定义了要访问的URLs，然后使用colly.NewCollector方法创建了一个默认的Collector实例。
- 接着，它使用redisstorage.Storage创建了一个Redis存储，并使用c.SetStorage方法将存储添加到Collector实例中。
- 使用存储可以方便地将数据存储到Redis数据库中，以便于后续处理和分析。
- 然后，程序清空了之前的存储数据，使用queue.New方法创建了一个基于Redis存储的请求队列。
- 在这个示例中，使用队列可以控制请求的并发量，避免因为请求过多导致服务器负荷过大。
- 接着，程序使用c.OnResponse方法设置回调函数，用于在请求完成后处理响应数据。
- 最后，程序将URLs添加到队列中，并使用q.Run方法开始请求队列中的请求。

这是一个把数据存储到redis的demo，如果为了搜索方便会把数据放到es中，其他更多储存操作会在`17_storage_backend`中介绍
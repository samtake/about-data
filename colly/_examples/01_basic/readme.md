### 最基本的使用流程
- 1.创建收集器：`colly.NewCollector`
- 2.监听回调方法,并处理业务逻辑：
    - OnRequest
    - OnResponse
    - OnHTML
    - OnXML
    - OnHTMLDetach
    - 其他 
- 3.开始访问:`c.Visit`

### OnHTML
OnHTML用于在HTML文档中选择元素并执行回调函数。它的定义：
``func (c *Collector) OnHTML(goquerySelector string, f HTMLCallback)``

其中，goquerySelector参数是一个字符串，用于选择HTML元素。它使用GoQuery Selectors语法，即Go语言的CSS选择器语法。这意味着，我们可以使用CSS选择器来选择HTML元素，并将它们传递给回调函数。



- `div.some-class`选择具有类名为some-class的div元素
- `#some-id` 选择具有ID为some-id的元素
- `a[href]`表示选择具有href属性的a元素
- `img[src='logo.png']`表示选择具有src属性值为logo.png的img元素
- `input[name^='user']`表示选择具有name属性以user开头的input元素
- `input[name$='pass']`表示选择具有name属性以pass结尾的input元素
- `a[href][class='external']`多个属性选择：选择具有href属性和类名为external的a元素
- `input[type='password'][name='username']`表示选择具有type属性为password和name属性为username的input元素

    注意，属性选择器的顺序很重要，因为它们必须与选择器中的顺序匹配。因此，在上面的示例中，如果我们交换type和name属性，则选择器将不匹配任何元素

- `a[href] > [class='external']`
    选择具有href属性的a元素的直接子元素，并且这些子元素具有class属性值为external。这个选择器中的>符号表示选择直接子元素。

    ```bash
    <div class="wrapper">
        <a href="https://www.example.com" target="_blank">
            <img src="example.png" class="external">
        </a>
    </div>
    ```
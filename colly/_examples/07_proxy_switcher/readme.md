### ProxySwitcher
ProxySwitcher是Colly爬虫框架中的一个接口，用于实现代理的切换。在Colly中，通过实现ProxySwitcher接口来实现代理的轮换或切换。
使用ProxySwitcher时，需要实现NextProxy方法，该方法会在每次请求时被调用，并返回下一个要使用的代理地址。在实现NextProxy方法时，可以根据不同的策略来选择下一个代理地址，例如轮询、随机等。
实现好ProxySwitcher接口后，可以通过SetProxyFunc方法将其设置为Collector的代理函数，这样在每次请求时就会自动调用NextProxy方法来选择代理地址。
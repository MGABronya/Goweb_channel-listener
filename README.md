# Goweb_channel-listener

用select关键字创建多通道监听器，先开启一个goroutine，然后用select关键字来监视各个通道数据输出并收集数据到通道。
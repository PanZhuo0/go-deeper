##### TCP网络编程 第4层

##### UDP网络编程 第4层

##### HTTP网络编程 第7层

##### web socket网络编程 第7层
    将HTTP升级即可变为websocket协议(协议头会被改变,将会建立流连接)
    与HTTP不同，websocket提供全双工通信，此外还支持基于TCP的消息流
    (HTTP请求通过后数据可以直接通过TCP传输,TCP read,write,不用再传那些HTTP header信息了)

##### Comet概念 
    Comet通信技术:是一种基于HTTP长连接的服务器推送数据的技术,感觉更像一种概念，不是具体的某项技术


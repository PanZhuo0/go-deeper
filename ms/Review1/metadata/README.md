Reference:
    https://grpc.io/docs/guides/metadata/

### 为什么在GRPC中需要Metadata
    在传统的请求调用中，存在这样的情况，用户发送请求，将认证信息(Token)/其他信息存放在请求头中,从而让服务端可以解析来确定请求的一些信息
    由于某些原因,这些信息直接放在请求Body中可能是不合适的,我们将这些信息放在HTTP Header中,而在RPC中MetaData就实现了类似HTTP Header的功能
### MetaData 的类型定义
```go
    type MD map[string][]string
```
MetaData底层式一个Map,Key为string  Value为[]string 类型

需要注意的点:
    1.key 不允许以grpc开头,grpc是GRPC协议保留的Metadata字段
    2.metadata中的key 不区分大小写
    3.对于二进制的值(比如图像、音频、文件),key需要以 `-bin` 结尾,在网络传输过程中，会对数据进行base64编码，到达后进行base64解码
### MetaData 的实现
    实际上，MetaData被附加到Client端/Server端 的Context上

### MetaData 的使用
    客户端发送MetaData
    服务端接收MetaData


    服务端发送MetaData
    客户端接收MetaData(客户端可以接收Header/ Trailer的MetaData)

##### 使用到Traier的情况
    1.需要计算通过流式传输给client的所有数据的MD5值,这个值需要传递完所有数据后，才能计算,需要使用到trailer MetaData


    Reference: 
        https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Trailer

### 普通RPC的Header/Trailer 中的Metadata使用

### 流式RPC中的Header/Trailer 中的MetaData使用
MetaData 有流式/普通  服务端接收/发送 客户端接收/发送
对于流式RPC，需要从Stream中获取Context,普通的则可以通过context获取MetaData

# MetaData示例
首先MetaData以发送方接收方分类，可以分为
    1.客户端发送、服务端接收MetaData
    2.服务端发送、客户端接收MetaData

    按照RPC的种类分类，可以分为
    1.流式RPC MetaData
    2.普通RPC MetaData

    按照MetaData本身,可以分为
    1.Header MetaData 请求头
    2.Trailer MetaData  请求尾

总体来说，MetaData 与HTTP Header/Trailer 对标


`示例:一个普通RPC、客户端发送并接收、服务端发送(带Header、Trailer的Metadata)并接收 示例`
客户端
    会发送一个MD:携带访问的token
    会接收从server端响应的Header、Trailer MD

服务端
    会接收客户端发来的MD 并对其进行解析、判断是否有效
    会响应给客户端Header MD(包含一个location key),Trailer MD(包含timestamp key)

Attention:服务端只应该使用从客户端发来的context,不能新建

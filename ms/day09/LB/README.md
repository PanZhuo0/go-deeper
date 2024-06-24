# 负载均衡LB 
这个示例使用在本机的多个端口，来模拟一个服务的不同节点

localhost:8976 1 
localhost:8977 2
localhost:8978 3
localhost:8979 4

在这里在consul中注册了4个实例，等着hello_cliet 发送RPC请求
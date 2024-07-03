### 错误处理
Reference: 
    https://grpc.io/docs/guides/status-codes/
    https://grpc.io/docs/guides/error/

1.gRPC中的错误类型
2.gRPC中往错误中追加信息
3.把错误信息中的追加信息解析出来
4.grpc.WithDetails 追加信息

### gRPC code

### gRPC status

# 错误处理示例
该Demo中将会限制Requst的请求频率,频率过高将会报错



server端给错误增加信息(status的 WithDetail)
```go

	if ok {
		// status 基本错误信息
		st := status.New(codes.ResourceExhausted, "request limit")
		// 往 status 中增加错误详细信息
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{
					{
						Subject:     fmt.Sprintf("name:%s", req.GetName()),
						Description: "每个Name 只能调用一次",
					},
				},
			},
		)
		if err != nil {
			fmt.Println(err)
			return nil, st.Err() //不携带额外信息的error
		} else {
			fmt.Println(ds.Details()...)
			return nil, ds.Err() //详细的error
		}
	} 

```


client端错误处理
```go

	resp, err := c.SayHello(context.Background(), &e.Request{
		Name: "张三",
	})
	if err != nil {
		// 解析错误中的详细信息(可能错误中携带了详细信息)
		s := status.Convert(err)
		ds := s.Details()
		for _, d := range ds {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Println("QuotaFailure: ", info)
			default:
				fmt.Println("unexpected type:", info)
			}
		}
		return
	}

```

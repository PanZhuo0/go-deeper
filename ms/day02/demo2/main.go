package main

import (
	"d2/demo2"
	"fmt"
)

func main() {
	// 客户端
	req := demo2.NoticeRequest{
		Msg:       "Email ",
		NoticeWay: &demo2.NoticeRequest_Email{Email: "xxx@xxx.com"},
	}

	req2 := demo2.NoticeRequest{
		Msg:       "Phone",
		NoticeWay: &demo2.NoticeRequest_Phone{Phone: "110"},
	}

	// Server 端
	// 类型断言
	switch v := req.NoticeWay.(type) {
	case *demo2.NoticeRequest_Email:
		fmt.Println("notice by email", v.Email)
	case *demo2.NoticeRequest_Phone:
		fmt.Println("notice by phone", v.Phone)
	}

	switch v2 := req2.NoticeWay.(type) {
	case *demo2.NoticeRequest_Email:
		fmt.Println("notice by email", v2.Email)
	case *demo2.NoticeRequest_Phone:
		fmt.Println("notice by phone", v2.Phone)
	}
}

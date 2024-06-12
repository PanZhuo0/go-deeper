package main

import (
	"d4/d4"
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/generator" //deprecated 过时了，没办法
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// 只有参考意义，kit库过时了
// 总之fieldmask可以实现仅传递消息中的部分信息,实现起来有点麻烦
func main() {
	// client
	paths := []string{"price", "info.date"} //更新的字段的信息

	r := d4.UpdateBook{
		Op: "operator_zhangsan",
		Book: &d4.Book{
			Price: &wrapperspb.Int64Value{Value: 45},
			Info: &d4.Book_Info{
				Data: "3024",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}

	//server
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(r.UpdateMask, generator.CamelCase) //这里不会超纲了
	// fielemask_util 支持读取到结构体等中，更多用法请看文档
	bookDst := make(map[string]interface{})
	fieldmask_utils.StructToMap(mask, r.Book, bookDst)
	fmt.Println("bookDst:", bookDst)
}

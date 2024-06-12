package main

import (
	"d3/d3"

	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	paths := []string{"price", "book.info.date"} //更新的字段的信息

	u := d3.UpdateBook{
		// 操作者
		Op: "operator_zhangsan",
		// 修改的书
		Book: &d3.Book{
			Price: &wrapperspb.Int64Value{Value: 45},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}

}

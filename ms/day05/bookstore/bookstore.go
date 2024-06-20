package main

import (
	"bookstore/pb"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedBookstoreServer
	bs *bookstore //data.go
}

func (s *server) ListShelves(ctx context.Context, in *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	// 调用ORM中的操作方法
	sl, err := s.bs.ListShelves(ctx)
	if err == gorm.ErrEmptySlice {
		// if Empty
		return &pb.ListShelvesResponse{}, nil
	}
	if err != nil { //retrive DB failed
		return nil, status.Error(codes.Internal, "Query DB failed")
	}
	// 返回数据
	nsl := make([]*pb.Shelf, 0, len(sl))
	for _, s := range sl {
		nsl = append(nsl, &pb.Shelf{
			Id:    s.ID,
			Theme: s.Theme,
			Size:  s.Size,
		})
	}
	return &pb.ListShelvesResponse{Shelves: nsl}, err
}

// 创建书架
func (s *server) CreateShelf(ctx context.Context, in *pb.CreateShelfRequest) (*pb.Shelf, error) {
	// 将从grpc请求中的信息解析到Shelf结构体中
	inData := Shelf{
		ID:    in.Shelf.Id,
		Theme: in.Shelf.Theme,
		Size:  in.Shelf.Size,
	}

	// 检查grpc请求中是否指定了书架的主题?
	if len(in.GetShelf().Theme) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid theme")
	}

	// 使用GORM创建数据
	ns, err := s.bs.CreateShelf(ctx, inData)
	if err != nil {
		return nil, status.Error(codes.Internal, "create failed")
	}

	// 将数据返回给grpc client
	return &pb.Shelf{
		Id:    ns.ID,
		Theme: ns.Theme,
		Size:  ns.Size,
	}, nil
}

/*
根据rpc Client 传入的书架ID获取书架信息

parameter:

	ctx context.Cotnext 	请求的上下文信息
	in pb.GetShelfRequest	grpc请求参数,内含书架ID

response:

	*pb.Shelf				返回对应书架ID的书架信息
	error					执行过程中可能出现的错误
*/
func (s *server) GetShelf(ctx context.Context, in *pb.GetShelfRequest) (*pb.Shelf, error) {
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	// 从GORM获取Shelf信息
	shelf, err := s.bs.GetShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}
	return &pb.Shelf{
		Id:    shelf.ID,
		Theme: shelf.Theme,
		Size:  shelf.Size,
	}, nil
}

/*
根据grpc client 传入的ID  删除书架

Parameter:

	ctx context.Context 			请求的上下文信息
	in *pb.DeleteShelfRequest		grpc client请求，包含了书架ID信息

Response:

	*emptypb.Empty
	error							可能出现的错误信息
*/
func (s *server) DeleteShelf(ctx context.Context, in *pb.DeleteShelfRequest) (*emptypb.Empty, error) {
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	// 调用GORM删除书架
	err := s.bs.DeleteShelv(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "delete failed")
	}
	return &empty.Empty{}, nil
}

/*
ListBooks 根据用户传入的参数列出所有的书籍信息

paramater:

	ctx context.Context					请求的上下文
	in *pb.ListBooksRequest				grpc请求，包括请求的分页参数

response:

	*pb.ListBooksResponse				grpc响应：包括了对应的图书集合、以及下一个分页游标token
*/
func (s *server) ListBooks(ctx context.Context, in *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	if in.GetPageToken() == "" {
		// return first page
	} else {
		pageInfo := Token(in.GetPageToken()).Decode()
		if pageInfo.IsValid() {
			return nil, status.Error(codes.InvalidArgument, "invalid page token")
		}
	}

	// 从GORM获取数据
	s.bs.GetBookListByShelfID(ctx, in.GetShelf(), 1)
	return nil, nil
}

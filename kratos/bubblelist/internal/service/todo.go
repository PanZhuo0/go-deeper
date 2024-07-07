package service

import (
	"context"
	"errors"

	pb "bubblelist/api/bubblelist/v1"
	v1 "bubblelist/api/bubblelist/v1"
	"bubblelist/internal/biz"
)

type TodoService struct {
	pb.UnimplementedTodoServer
	// 嵌入一个实现业务逻辑的结构体(biz层)
	uc *biz.TodoUsecase
}

func NewTodoService(uc *biz.TodoUsecase) *TodoService {
	return &TodoService{
		uc: uc,
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	// 1.参数校验
	if len(req.GetTitle()) == 0 {
		return &pb.CreateTodoReply{}, errors.New("无效的title")
	}
	// 2.调用业务逻辑
	data, err := s.uc.Create(ctx, &biz.Todo{Title: req.Title})
	if err != nil {
		return nil, err
	}
	// 3.返回响应
	return &pb.CreateTodoReply{
		Id:     data.ID,
		Title:  data.Title,
		Status: data.Status,
	}, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoReply, error) {
	// 1.参数校验
	if req.Id <= 0 {
		return &pb.UpdateTodoReply{}, errors.New("无效的ID")
	}
	// 2.调用biz层业务逻辑
	err := s.uc.Update(ctx, &biz.Todo{ID: req.Id, Title: req.Title, Status: req.Status})
	// 3.响应
	return &pb.UpdateTodoReply{}, err
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	// 1.参数校验
	if req.Id <= 0 {
		return &pb.DeleteTodoReply{}, errors.New("无效的ID")
	}
	// 2.调用biz层业务逻辑
	err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// 3.响应结果
	return nil, err
}

func (s *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoReply, error) {
	// 1.参数校验
	if req.Id <= 0 {
		return &pb.GetTodoReply{}, errors.New("无效的ID")
	}
	// 2.调用biz层业务逻辑
	ret, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		// return nil, err
		// return 自定义的错误
		return nil, v1.ErrorTodoNotFound("id:%v todo is not found", req.Id)
	}
	// 3.响应
	return &pb.GetTodoReply{
		Id:     ret.ID,
		Title:  ret.Title,
		Status: ret.Status,
	}, nil
}

func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	return &pb.ListTodoReply{}, nil
}

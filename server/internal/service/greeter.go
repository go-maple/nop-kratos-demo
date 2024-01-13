package service

import (
	"context"

	"github.com/go-maple/nop-kratos-demo/server/internal/biz"

	v1 "github.com/go-maple/nop-kratos-demo/server/api/server/v1"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.SayHello(ctx, &biz.Greeter{Name: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Name}, nil
}

func (s *GreeterService) SayHelloError(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return nil, v1.ErrorCreateServerFailed("SayHelloError")
}

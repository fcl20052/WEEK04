package service



import (
	"context"
	"errors"
	v1 "liquid/api/helloworld/v1"
	"liquid/internal/biz"
)



type GreeterService struct {
	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if in.GetName() == "error" {
		return nil, errors.New("user not found")
	}
	return &v1.HelloReply{Message: "Hello" + in.GetName()}, nil
}

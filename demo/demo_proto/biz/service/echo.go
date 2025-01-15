package service

import (
	"context"

	pbapi "github.com/North-al/hertz_demo/kitex_gen/pbapi"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {
	// 直接返回请求中的消息
	return &pbapi.Response{
		Message: req.Message,
	}, nil
}

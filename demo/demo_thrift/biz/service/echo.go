package service

import (
	"context"
	"fmt"

	api "github.com/North-al/hertz_demo/demo/demo_thrift/kitex_gen/api"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(request *api.Request) (resp *api.Response, err error) {
	// 直接返回请求中的消息

	info := rpcinfo.GetRPCInfo(s.ctx)
	fmt.Printf("info %v\n", info.From().ServiceName())

	return &api.Response{
		Message: request.Message,
	}, nil
}

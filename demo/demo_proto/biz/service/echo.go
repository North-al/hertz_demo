package service

import (
	"context"
	"fmt"

	pbapi "github.com/North-al/hertz_demo/kitex_gen/pbapi"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {
	clientName, ok := metainfo.GetPersistentValue(s.ctx, "CLIENT_NAME")
	fmt.Println("clientName: ", clientName, "ok: ", ok)

	err = kerrors.NewBizStatusError(404, "not found")
	return nil, err

	// return &pbapi.Response{
	// 	Message: req.Message,
	// }, nil
}

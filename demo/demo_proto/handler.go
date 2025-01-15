package main

import (
	"context"
	pbapi "github.com/North-al/hertz_demo/kitex_gen/pbapi"
	"github.com/North-al/hertz_demo/biz/service"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *pbapi.Request) (resp *pbapi.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}

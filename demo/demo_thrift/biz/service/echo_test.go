package service

import (
	"context"
	api "github.com/North-al/hertz_demo/demo/demo_thrift/kitex_gen/api"
	"testing"
)

func TestEcho_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoService(ctx)
	// init req and assert value

	request := &api.Request{}
	resp, err := s.Run(request)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

package service

import (
	"context"

	common "github.com/North-al/hertz_demo/mall/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SignOutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSignOutService(Context context.Context, RequestContext *app.RequestContext) *SignOutService {
	return &SignOutService{RequestContext: RequestContext, Context: Context}
}

func (h *SignOutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	session := sessions.Default(h.RequestContext)
	session.Clear()
	err = session.Save()
	if err != nil {
		return nil, err
	}

	return
}

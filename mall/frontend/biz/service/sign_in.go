package service

import (
	"context"

	auth "github.com/North-al/hertz_demo/mall/frontend/hertz_gen/frontend/auth"
	common "github.com/North-al/hertz_demo/mall/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SignInService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSignInService(Context context.Context, RequestContext *app.RequestContext) *SignInService {
	return &SignInService{RequestContext: RequestContext, Context: Context}
}

func (h *SignInService) Run(req *auth.LoginRequest) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// 登录

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}

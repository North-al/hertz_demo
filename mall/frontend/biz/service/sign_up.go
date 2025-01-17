package service

import (
	"context"

	auth "github.com/North-al/hertz_demo/mall/frontend/hertz_gen/frontend/auth"
	common "github.com/North-al/hertz_demo/mall/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type SignUpService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSignUpService(Context context.Context, RequestContext *app.RequestContext) *SignUpService {
	return &SignUpService{RequestContext: RequestContext, Context: Context}
}

func (h *SignUpService) Run(req *auth.RegisterRequest) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// 注册
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	return
}

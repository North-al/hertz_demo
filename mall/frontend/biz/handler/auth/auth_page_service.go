package auth

import (
	"context"

	"github.com/North-al/hertz_demo/mall/frontend/biz/service"
	"github.com/North-al/hertz_demo/mall/frontend/biz/utils"
	auth "github.com/North-al/hertz_demo/mall/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SignIn .
// @router /auth/sign-in [POST]
func SignIn(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	_, err = service.NewSignInService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, "111")
	c.Redirect(consts.StatusOK, []byte("/"))
}

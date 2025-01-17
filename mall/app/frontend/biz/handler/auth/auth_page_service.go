package auth

import (
	"context"
	"fmt"

	"github.com/North-al/hertz_demo/mall/app/frontend/biz/service"
	"github.com/North-al/hertz_demo/mall/app/frontend/biz/utils"
	auth "github.com/North-al/hertz_demo/mall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/North-al/hertz_demo/mall/app/frontend/hertz_gen/frontend/common"
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
	redirectUrl, err := service.NewSignInService(ctx, c).Run(&req)
	fmt.Println("redirectUrl", redirectUrl)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, "111")
	c.Redirect(consts.StatusOK, []byte(redirectUrl))
}

// SignUp .
// @router /auth/sign-up [POST]
func SignUp(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewSignUpService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, "000")

	c.Redirect(consts.StatusOK, []byte("/sign-in"))
}

// SignOut .
// @router /auth/sign-out [GET]
func SignOut(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewSignOutService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Redirect(consts.StatusOK, []byte("/"))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

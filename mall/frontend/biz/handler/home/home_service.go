package home

import (
	"context"

	"github.com/North-al/hertz_demo/mall/frontend/biz/service"
	"github.com/North-al/hertz_demo/mall/frontend/biz/utils"
	"github.com/North-al/hertz_demo/mall/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	// resp := []map[string]interface{}{}
	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)

	// resp["user_id"] = 1

	c.HTML(consts.StatusOK, "home.tmpl", utils.WrapResponse(ctx, c, resp))

}

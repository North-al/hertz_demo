package service

import (
	"context"

	"github.com/North-al/hertz_demo/mall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (resp map[string]interface{}, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp = make(map[string]interface{})

	items := []map[string]interface{}{
		{
			"title":   "T-shirt",
			"price":   100,
			"picture": "/static/images/t-shirt.webp",
		},
		{
			"title":   "Short sleeve",
			"price":   120,
			"picture": "/static/images/t-shirt.webp",
		},
		{
			"title":   "Long sleeve",
			"price":   150,
			"picture": "/static/images/t-shirt.webp",
		},
		{
			"title":   "Hoodie",
			"price":   180,
			"picture": "/static/images/t-shirt.webp",
		},
		{
			"title":   "avatar",
			"price":   10,
			"picture": "/static/images/avatar.jpg",
		},
	}

	resp["items"] = items

	return resp, nil
}

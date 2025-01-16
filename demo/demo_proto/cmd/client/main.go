package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/North-al/hertz_demo/kitex_gen/pbapi"
	"github.com/North-al/hertz_demo/kitex_gen/pbapi/echo"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatalf("创建 Consul 解析器失败: %v", err)
	}

	c, err := echo.NewClient(
		"demo_proto",
		client.WithResolver(r),
		// client.WithTransportProtocol(transport.GRPC), // window11中grpc支持有问题
		client.WithShortConnection(), // 使用短链接
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	if err != nil {
		log.Fatalf("创建 Echo 客户端失败: %v", err)
	}

	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	res, err := c.Echo(ctx, &pbapi.Request{Message: "error"}) // 传递错误消息

	var bizErr *kerrors.BizStatusError // 修改为通用的 BizStatusError
	fmt.Println("err: ", err)
	fmt.Println("res: ", res)

	if err != nil {
		if errors.As(err, &bizErr) {
			fmt.Printf("捕获到业务错误: 错误码=%d, 错误信息=%s\n", bizErr.BizStatusCode(), bizErr.BizMessage())
			return
		}
		log.Fatalf("调用 Echo 服务失败: %v", err)
	}

	if res == nil {
		log.Fatalln("服务器响应为空")
	} else {
		fmt.Printf("服务器响应: %v\n", res)
	}
}

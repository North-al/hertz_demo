package main

import (
	"context"
	"fmt"
	"log"

	"github.com/North-al/hertz_demo/kitex_gen/pbapi"
	"github.com/North-al/hertz_demo/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatalf("创建 Consul 解析器失败: %v", err)
	}

	c, err := echo.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		log.Fatalf("创建 Echo 客户端失败: %v", err)
	}

	fmt.Printf("client %v\n", c)

	res, err := c.Echo(context.TODO(), &pbapi.Request{
		Message: "Hello Server",
	})
	if err != nil {
		log.Fatalf("调用 Echo 方法失败: %v", err)
	}

	fmt.Printf("服务器响应: %v\n", res)
}

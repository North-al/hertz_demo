package dal

import (
	"github.com/North-al/hertz_demo/demo/demo_thrift/biz/dal/mysql"
	"github.com/North-al/hertz_demo/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

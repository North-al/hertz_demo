package dal

import (
	"github.com/North-al/hertz_demo/mall/app/frontend/biz/dal/mysql"
	"github.com/North-al/hertz_demo/mall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

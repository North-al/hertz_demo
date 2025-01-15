package dal

import (
	"github.com/North-al/hertz_demo/biz/dal/mysql"
	"github.com/North-al/hertz_demo/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

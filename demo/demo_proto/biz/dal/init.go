package dal

import (
	"github.com/North-al/hertz_demo/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}

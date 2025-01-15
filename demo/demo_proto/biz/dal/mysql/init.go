package mysql

import (
	"fmt"
	"os"

	"github.com/North-al/hertz_demo/biz/model"
	"github.com/North-al/hertz_demo/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	fmt.Printf("conf %v\n", conf.GetConf())
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DB"),
	)

	fmt.Printf("dsn %v\n", dsn)

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.User{})

}

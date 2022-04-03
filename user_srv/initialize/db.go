package initialize

import (
	"fmt"
	"github.com/luhaopei/mxshop_srvs/user_srv/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.ServerConfig.MySqlInfo.User,
		global.ServerConfig.MySqlInfo.Password,
		global.ServerConfig.MySqlInfo.Host,
		global.ServerConfig.MySqlInfo.Port,
		global.ServerConfig.MySqlInfo.Name,
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{

			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/luhaopei/mxshop_srvs/goods_srv/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

// 随机字符串+用户密码
func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{

			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = db.AutoMigrate(&model.Category{}, &model.Brands{},
		&model.GoodsCategoryBrand{}, &model.Banner{},
		&model.Goods{})

}

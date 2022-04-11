package global

import (
	"github.com/luhaopei/mxshop_srvs/goods_srv/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)

func init() {

}

package global

import (
	"github.com/luhaopei/mxshop_srvs/user_srv/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)

func init() {

}

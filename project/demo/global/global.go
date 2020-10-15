package global

import (
	"demo/pkg/config"
	"github.com/jinzhu/gorm"
)

var ConfigSetting config.Configure
var ServerConfig config.Configure
var AppConfig config.Configure
var DataBaseConfig config.Configure
var DBEngine *gorm.DB

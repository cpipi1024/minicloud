package bootstrap

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SrvConf   GlobalConf                                     // 全局配置
	MSDB      *gorm.DB                                       // mysqldb
	Logger    *zap.Logger                                    // Logger
	confPath  = "/home/yangtao/workspace/minicloud/app.yaml" // 配置路径
	injectors []injecotor                                    // mysql依赖
)

// todo: 加载全局配置
func LoadConf() {
	v := viper.New()

	v.SetConfigFile(confPath)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := GlobalConf{}

	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}

	SrvConf = conf

}

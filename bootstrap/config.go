package bootstrap

// 全局配置
type GlobalConf struct {
	App       AppConifg   `mapstructure:"app"`
	Mysql     MySQLConfig `mapstructure:"mysql"`
	JWT       JWTConfig   `mapstructure:"jwt"`
	Log       LogConfig   `mapstructure:"log"`
	CloudPath string      `mapstructure:"cloudpath"`
}

// app配置
type AppConifg struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// mysql配置
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

// jwt配置
type JWTConfig struct {
	SecretKey string `mapstructure:"secretkey"` // 私钥
	TTL       int    `mapstructure:"ttl"`       // ttl
}

// zaplogger配置
type LogConfig struct {
	Level string `mapstructure:"level"` // 日志级别
	Path  string `mapstructure:"path"`  // 日志路径
}

package bootstrap

import "go.uber.org/zap"

// todo: 创建logger
func LoadLogger() {
	logconf := SrvConf.Log

	var zapconf zap.Config

	switch logconf.Level {
	// 开发环境
	case "dev":
		zapconf = zap.NewDevelopmentConfig()

	// 生产环境
	case "pro":
		zapconf = zap.NewProductionConfig()
		zapconf.OutputPaths = append(zapconf.OutputPaths, logconf.Path)
		zapconf.ErrorOutputPaths = append(zapconf.ErrorOutputPaths, logconf.Path)
	}

	logger, err := zapconf.Build()

	if err != nil {
		panic(err)
	}

	Logger = logger

}

// todo: 接口error日志
func ApiError(info string, method string, path string, code int, err string) {
	Logger.Error(
		info,
		zap.String("mthod", method),
		zap.String("path", path),
		zap.Int("code", code),
		zap.String("err", err),
	)
}

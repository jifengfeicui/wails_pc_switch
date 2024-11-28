package global

import (
	"context"
	"os"

	"go.uber.org/zap"

	"naive_test2/config"
)

//var DB *gorm.DB

var (
	Dir         string
	SugarLogger *zap.SugaredLogger
	StartupCtx  context.Context
	CONFIG      config.System // 系统配置信息
)

func init() {
	Dir, _ = os.Getwd()

}

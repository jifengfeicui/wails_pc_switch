package initialize

import (
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"naive_test2/global"
)

func Viper() {
	// 设置配置文件类型和路径
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")
	// 读取配置信息
	err := viper.ReadInConfig()
	dir, _ := os.Getwd()
	global.SugarLogger.Info("当前工作目录：", dir)
	if err != nil {
		global.SugarLogger.Panic("获取配置文件错误")
	}
	// 将读取到的配置信息反序列化到全局 CONFIG 中
	err = viper.Unmarshal(&global.CONFIG)
	if err != nil {
		global.SugarLogger.Panic("viper反序列化错误")
	}
	//global.SugarLogger.Debug(global.CONFIG)
	// 监视配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		global.SugarLogger.Debug("配置文件被修改：", e.Name)
	})
}

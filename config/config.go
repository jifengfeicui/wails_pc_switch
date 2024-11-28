package config

// System 定义项目配置文件结构体
type System struct {
	BafaConfig BafaConfig `mapstructure:"bafa"`
}

// GinConfig 定义 Gin 配置文件的结构体
type BafaConfig struct {
	Uid   string `mapstructure:"uid"`
	Topic string `mapstructure:"topic"`
}

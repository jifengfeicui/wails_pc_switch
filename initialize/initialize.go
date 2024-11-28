package initialize

func Initialize() {
	//初始化方法,调试模式勿打开 ChangeWorkingDir
	//ChangeWorkingDir()
	CreateMkdirall()
	InitLogger()
	Viper()
}

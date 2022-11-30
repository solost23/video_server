package initialize

func Initialize(configPath string) {
	InitConfig(configPath)
	InitMysql()
	InitOSSClient()
}

package initialize

func Initialize(configPath string) {
	InitConfig(configPath)
	InitMysql(false)
	InitMysql(true)
	InitOSSClient()
}

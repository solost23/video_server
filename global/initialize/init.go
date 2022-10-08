package initialize

func Initialize() {
	InitConfig("./config/config.yml")
	InitMysql(false)
	InitMysql(true)
	InitMinio()
}

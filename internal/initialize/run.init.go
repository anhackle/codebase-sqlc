package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysqlC()
	InitRedis()
	InitServiceInterface()
	InitValidator()

	r := InitRouter()
	r.Run(":8082")

}

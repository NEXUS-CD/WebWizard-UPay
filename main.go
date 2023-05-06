package main

import (
	"io"
	"log"
	"os"

	"UPay/configs"
	"UPay/database"
	"UPay/logger"
	"UPay/router"
)

func main() {
	configs.GetConfig()
	// 初始化日志记录器
	logger.InitLogger()
	logFile, err := os.OpenFile("logger/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	// 初始化数据库连接
	database.InitMongoDB()
	database.InitRedis()
	// 初始化路由
	r := router.SetupRouter()
	// 启动服务器
	port := configs.GetConfig().Server.Port
	log.Printf("Server is running at http://localhost:%s\n", port)
	r.Run(":" + port)
}

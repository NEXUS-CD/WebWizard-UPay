package logger

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger() {
	log = logrus.New()

	// 设置日志输出级别
	log.SetLevel(logrus.InfoLevel)

	// 设置日志输出格式
	formatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetFormatter(formatter)

	// 设置日志输出目标，默认为控制台
	log.SetOutput(os.Stdout)

	// 如果需要将日志记录到文件，可以取消下面代码的注释并替换为实际的文件路径
	logFile, err := os.OpenFile("logger/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

}

func GetLogger() *logrus.Logger {
	return log
}

// LoggerMiddleware 记录请求日志的中间件函数
func LoggerMiddleware() gin.HandlerFunc {
	// 因为加了中间件，进行输出日志+终端输出，导致终端输出两次，需解决
	return func(c *gin.Context) {
		start := time.Now()
		// 处理请求
		c.Next()
		// 记录请求日志
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		latency := time.Since(start)
		log.Printf("[+] %3d | %13v | %15s | %-7s %s\n",
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}

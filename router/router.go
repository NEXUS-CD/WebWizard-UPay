package router

import (
	"github.com/NEXUS-CD/UPay/controllers"
	"github.com/NEXUS-CD/UPay/logger"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 添加中间件监听路由请求，并记录日志
	r.Use(logger.LoggerMiddleware())
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id", controllers.GetUser)
	}

	return r
}

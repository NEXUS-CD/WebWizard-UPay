package router

import (
	"UPay/controllers"
	"UPay/docs"
	"UPay/logger"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 添加中间件监听路由请求，并记录日志
	r.Use(logger.LoggerMiddleware())
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id", controllers.GetUser)
	}

	return r
}

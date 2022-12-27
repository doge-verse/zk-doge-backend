package api

import (
	"github.com/doge-verse/zk-doge-backend/api/handler"
	"github.com/doge-verse/zk-doge-backend/api/middleware"
	"github.com/doge-verse/zk-doge-backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors())

	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Router := router.Group("/api/v1")
	initNoAuthRouter(v1Router)
	initNeedAuthRouter(v1Router)
}

func initNoAuthRouter(r *gin.RouterGroup) {
	r.POST("/user/login", handler.Login)
	r.POST("/user/register", handler.Register)
}

func initNeedAuthRouter(r *gin.RouterGroup) {
	r.Use(middleware.JWTAuth())
}

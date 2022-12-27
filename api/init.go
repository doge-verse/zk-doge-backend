package api

import (
	"github.com/doge-verse/zk-doge-backend/pkg/conf"
	"github.com/gin-gonic/gin"
)

// Init .
func Init() *gin.Engine {
	gin.SetMode(conf.Cfg.Gin.Mode)
	r := gin.Default()
	InitRouter(r)
	return r
}
